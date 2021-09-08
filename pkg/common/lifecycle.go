package common

import (
	"fmt"
	"os"
	"os/exec"
)

type Lifecycle struct {
	RunCommandBefore *string `yaml:"runBefore" json:"runBefore"`
	RunCommandAfter  *string `yaml:"runAfter" json:"runAfter"`
}

func (h Lifecycle) runCommand(command string) error {
	cmd := exec.Command(command)
	fmt.Println(string(command))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func (h Lifecycle) HandleRunCommandBefore() error {
	if h.RunCommandBefore != nil {
		h.runCommand(*h.RunCommandBefore)
	}
	return nil
}

func (h Lifecycle) HandleRunCommandAfter() error {
	if h.RunCommandAfter != nil {
		h.runCommand(*h.RunCommandAfter)
	}
	return nil
}
