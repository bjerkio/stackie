package config

import (
	"fmt"
	"path"
)

func (config *Config) GetPath(basePath string, fileNames []FileName) (string, error) {
	for _, f := range fileNames {
		p := path.Join(basePath, f.FileName)
		_, err := config.fs.Stat(p)
		if err == nil {
			return p, nil
		}
	}

	return "", fmt.Errorf("could not find configuration")
}
