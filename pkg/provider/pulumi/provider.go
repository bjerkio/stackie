package pulumi

type PulumiProfileConfig struct {
	AccessToken string `yaml:"access-token"`
}

type PulumiProjectConfig struct {
	StackName string
	CloudURL string // A cloud URL to log in to
	AccessToken string `yaml:"accessToken"`
	ProfileName string `yaml:"profileName"`
}