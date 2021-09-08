package types

import "github.com/bjerkio/stackie/pkg/common"

type PulumiStack struct {
	StackName string `yaml:"stackName" json:"stackName"`
	CloudURL  string `yaml:"cloudUrl" json:"cloudUrl"`
}

type PulumiUserConfig struct {
	Name      string             `yaml:"name" json:"name"`
	Lifecycle []common.Lifecycle `yaml:"lifecycle" json:"lifecycle"`
}
