package config

import "github.com/spf13/afero"

type Config struct {
	fs afero.Fs
}

func NewConfig() Config {
	fs := afero.NewOsFs()
	c := Config{fs}
	return c
}

func (config *Config) SetFS(fs afero.Fs) {
	config.fs = fs
}
