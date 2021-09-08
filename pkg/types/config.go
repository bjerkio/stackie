package types

import (
	pulumiTypes "github.com/bjerkio/stackie/pkg/provider/pulumi/types"
)

type Configs struct {
	UserConfig        *UserConfig
	ProjectConfig     *ProjectConfig
	UserProjectConfig *UserProjectConfig
}

type UserConfig struct {
	Pulumi []pulumiTypes.PulumiUserConfig `yaml:"pulumi" json:"pulumi"`
}

type ProjectStack struct {
	Environment *string                  `yaml:"environment" json:"environment"`
	Pulumi      *pulumiTypes.PulumiStack `yaml:"pulumi" json:"pulumi"`
}

type ProjectConfig struct {
	Name   string                  `yaml:"name" json:"name"`
	Stacks map[string]ProjectStack `yaml:"stacks" json:"stacks"`
}

type UserProjectConfig struct {
	ActiveStack *string `yaml:"activeStack" json:"activeStack"`
}
