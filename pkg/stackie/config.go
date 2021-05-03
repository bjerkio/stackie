package stackie

import (
	"github.com/bjerkio/stackie/pkg/provider/pulumi"
)

type UserConfig struct {
	Pulumi []pulumi.PulumiUserConfig
}

type ProjectStack struct {
	Name        string
	Environment string
	Pulumi      pulumi.PulumiStack
}

type ProjectConfig struct {
	Name   string
	Stacks []ProjectStack
}

type UserProjectConfig struct {
	ActiveStack string
}
