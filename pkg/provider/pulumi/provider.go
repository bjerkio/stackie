package pulumi

type PulumiStack struct {
	StackName string
	CloudURL  string
}

type PulumiUserProfile struct {
	Name        string
	AccessToken string
}

type PulumiUserConfig struct {
	AccessToken string
	Profiles    []PulumiUserProfile
}
