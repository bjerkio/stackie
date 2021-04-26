package stackie

import (
	"os"
	"path"
	"io/ioutil"
	"gopkg.in/yaml.v3"
	"github.com/bjerkio/stackie/pkg/provider/pulumi"
	log "github.com/sirupsen/logrus"
)

type ProjectConfig struct {
	Pulumi pulumi.PulumiProjectConfig
}

func getProjectConfig() (*ProjectConfig, error) {
	curPath, err := os.Getwd()
	if err != nil {
		// log.
		return nil, err
	}
	
	configPath := path.Join(curPath, ".stackie.yml")
	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Info("No stackie profile config found in this directory.")
		return nil, err
	}

	var config ProjectConfig
	err = yaml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}

	log.Info(config)
	return nil, nil
}

// func 