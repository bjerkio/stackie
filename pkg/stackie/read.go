package stackie

import (
	"fmt"
	"os"
	"path"

	"github.com/bjerkio/stackie/pkg/config"
	"github.com/bjerkio/stackie/pkg/types"
)

const CONFIG_FOLDER = "stackie"
const CONFIG_NAME = "stackie.yml"

// e.g. ~/.config/stackie/stackie.yml
func readUserConfig(c config.Config) (*types.UserConfig, error) {
	var conf types.UserConfig
	homeDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	configPath := path.Join(homeDir, CONFIG_FOLDER, CONFIG_NAME)
	fmt.Println(configPath)
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
func readProjectConfig(c config.Config) (*types.ProjectConfig, error) {
	var conf types.ProjectConfig
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
func readUserProjectConfig(c config.Config) (*types.UserProjectConfig, error) {
	var conf types.UserProjectConfig
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
func ReadConfigs(c config.Config) (*types.Configs, error) {

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

	return &types.Configs{
		UserConfig,
		ProjectConfig,
		UserProjectConfig,
	}, nil
}
