package pulumi

import (
	"time"

	"github.com/pulumi/pulumi/pkg/v3/backend/state"
	"github.com/pulumi/pulumi/sdk/v3/go/common/workspace"
)

func getPulumiConfig() (*workspace.Credentials, error) {
	creds, err := workspace.GetStoredCredentials()
	if err != nil {
		return nil, err
	}

	return &creds, nil
}

func setPulumiConfig(current string) error {
	return workspace.StoreAccount(current, workspace.Account{LastValidatedAt: time.Now()}, true)
}

func setPulumiStack(stackName string) error {
	return state.SetCurrentStack(stackName)
}
