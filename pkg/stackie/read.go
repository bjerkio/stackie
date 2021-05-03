package stackie

import (
	"os"
	"path"

	"github.com/bjerkio/stackie/pkg/config"
)

const CONFIG_FOLDER = "stackie"
const CONFIG_NAME = "stackie.yml"

type Configs struct {
	UserConfig        *UserConfig
	ProjectConfig     *ProjectConfig
	UserProjectConfig *UserProjectConfig
}

// e.g. ~/.config/stackie/stackie.yml
func readUserConfig(c config.Config) (*UserConfig, error) {
	var conf UserConfig
	homeDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := path.Join(homeDir, CONFIG_FOLDER, CONFIG_NAME)
	if !c.Exists(configPath) {
		return nil, nil
	}

	err = c.Read(configPath, config.YAML, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// e.g. project/.stackie.yaml
func readProjectConfig(c config.Config) (*ProjectConfig, error) {
	var conf ProjectConfig
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	p, ft, err := c.GetPath(wd, config.GetProjectFileNames())
	if err != nil {
		return nil, nil
	}

	err = c.Read(p, ft.Type, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// e.g. project/.stackie.local.yaml
func readUserProjectConfig(c config.Config) (*UserProjectConfig, error) {
	var conf UserProjectConfig
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	p, ft, err := c.GetPath(wd, config.GetUserFileNames())
	if err != nil {
		return nil, nil
	}

	err = c.Read(p, ft.Type, &conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}

// ReadConfigs returns all possible configurations
func ReadConfigs(c config.Config) (*Configs, error) {

	UserConfig, err := readUserConfig(c)
	if err != nil {
		return nil, err
	}

	ProjectConfig, err := readProjectConfig(c)
	if err != nil {
		return nil, err
	}

	UserProjectConfig, err := readUserProjectConfig(c)
	if err != nil {
		return nil, err
	}

	confs := Configs{
		UserConfig,
		ProjectConfig,
		UserProjectConfig,
	}

	return &confs, nil
}
