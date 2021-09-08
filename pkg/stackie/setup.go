package stackie

import (
	"errors"

	"github.com/bjerkio/stackie/pkg/config"
	"github.com/bjerkio/stackie/pkg/provider/pulumi"
	"github.com/bjerkio/stackie/pkg/types"
)

func Setup() error {
	c := config.New()
	conf, err := ReadConfigs(c)
	if err != nil {
		return err
	}

	if conf.ProjectConfig.Stacks == nil {
		return nil
	}

	var activeStack types.ProjectStack

	// Get active stack
	if len(conf.ProjectConfig.Stacks) == 1 {
		for _, s := range conf.ProjectConfig.Stacks {
			activeStack = s
		}
	}

	if len(conf.ProjectConfig.Stacks) > 1 {
		if conf.UserProjectConfig.ActiveStack == nil {
			return errors.New("Missing active stack")
		}

		a := *conf.UserProjectConfig.ActiveStack
		activeStack = conf.ProjectConfig.Stacks[a]
	}

	// Pulumi Provider
	p := pulumi.New(conf, activeStack)
	err = p.Prepare()

	if err != nil {
		return err
	}

	return nil
}
