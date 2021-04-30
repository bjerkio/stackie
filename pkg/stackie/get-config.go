package stackie

import (
	"fmt"
	"path"

	"github.com/bjerkio/stackie/pkg/provider/pulumi"
	"github.com/spf13/afero"
)

type ProjectConfig struct {
	Pulumi pulumi.PulumiProjectConfig
}

func getConfigFileNames() []string {
	return []string{
		".stackie.yml",
		".stackie.json",
		"stackie.yml",
		"stackie.json",
	}
}

type ConfigFS struct {
	fs afero.Fs
}

func NewConfigFS() ConfigFS {
	fs := afero.NewOsFs()
	c := ConfigFS{fs}
	return c
}

func (config *ConfigFS) getPath(basePath string) (string, error) {
	configFileNames := getConfigFileNames()
	for _, f := range configFileNames {
		p := path.Join(basePath, f)
		_, err := config.fs.Stat(p)
		if err == nil {
			return p, nil
		}
	}

	return "", fmt.Errorf("could not find configuration")
}

// func (config *ConfigFS) read(filePath string) (string, error) {
// 	ext := filepath.Ext(filePath)

// 	conf := &ProjectConfig{}

// 	switch ext {
// 	case "yml":
// 	case "yaml":
// 		err = yaml.Unmarshal(c, conf)
// 		if err != nil {
// 			return "", err
// 		}
// 	case "json":
// 		err = json.Unmarshal(c., conf)
// 		if err != nil {
// 			return "", err
// 		}
// 	}
// }

// func getProjectConfig() (*ProjectConfig, error) {
// 	curPath, err := os.Getwd()
// 	if err != nil {
// 		// log.
// 		return nil, err
// 	}

// 	configPath := path.Join(curPath, ".stackie.yml")
// 	content, err := ioutil.ReadFile(configPath)
// 	if err != nil {
// 		log.Info("No stackie profile config found in this directory.")
// 		return nil, err
// 	}

// 	var config ProjectConfig
// 	err = yaml.Unmarshal(content, config)
// 	if err != nil {
// 		return nil, err
// 	}

// 	log.Info(config)
// 	return nil, nil
// }
