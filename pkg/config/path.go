package config

import (
	"fmt"
	"path"
)

func (config *Config) GetPath(basePath string, fileNames []FileName) (string, *FileName, error) {
	for _, f := range fileNames {
		p := path.Join(basePath, f.FileName)
		_, err := config.fs.Stat(p)
		if err == nil {
			return p, &f, nil
		}
	}

	return "", nil, fmt.Errorf("could not find configuration")
}
