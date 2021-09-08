package config

import (
	"encoding/json"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// Read Configuration file
func (config *Config) Read(configPath string, fileType FileType, out interface{}) error {
	content, err := afero.ReadFile(config.fs, configPath)
	if err != nil {
		return err
	}

	switch fileType {
	case YAML:
		err = yaml.Unmarshal(content, out)
	case JSON:
		err = json.Unmarshal(content, out)
	}

	if err != nil {
		return err
	}

	return nil
}
