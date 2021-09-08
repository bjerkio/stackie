package config

func (config *Config) Exists(basePath string) bool {
	_, err := config.fs.Stat(basePath)
	if err != nil {
		return false
	}
	return true
}
