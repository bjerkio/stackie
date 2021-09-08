package pulumi

import (
	"github.com/bjerkio/stackie/pkg/provider"
	pulumiTypes "github.com/bjerkio/stackie/pkg/provider/pulumi/types"
	"github.com/bjerkio/stackie/pkg/types"
)

type PulumiProvider struct {
	userConfig  *types.UserConfig
	project     *types.ProjectConfig
	userProject *types.UserProjectConfig
	activeStack types.ProjectStack
}

func (p PulumiProvider) getCurrentProfile() (*pulumiTypes.PulumiUserConfig, error) {
	for i, u := range p.userConfig.Pulumi {
		if p.project.Name == u.Name {
			return &p.userConfig.Pulumi[i], nil
		}
	}

	return nil, nil
}

func (p PulumiProvider) Prepare() error {
	if p.activeStack.Pulumi == nil {
		return nil
	}

	profile, err := p.getCurrentProfile()
	if err != nil {
		return err
	}

	if p.activeStack.Pulumi.CloudURL != "" {

		pc, err := getPulumiConfig()
		if err != nil {
			return err
		}

		if pc.Current == p.activeStack.Pulumi.CloudURL {
			return nil
		}

		setPulumiConfig(p.activeStack.Pulumi.CloudURL)
	}

	if profile != nil {
		for _, h := range profile.Lifecycle {
			err = h.HandleRunCommandBefore()
			if err != nil {
				return err
			}
		}

		err := setPulumiStack(p.activeStack.Pulumi.StackName)
		if err != nil {
			return err
		}

		for _, h := range profile.Lifecycle {
			err = h.HandleRunCommandAfter()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func New(c *types.Configs, activeStack types.ProjectStack) provider.StackieProvider {
	var p provider.StackieProvider = PulumiProvider{
		userConfig:  c.UserConfig,
		project:     c.ProjectConfig,
		userProject: c.UserProjectConfig,
		activeStack: activeStack,
	}
	return p
}
