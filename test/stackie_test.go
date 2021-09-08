package test

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/bjerkio/stackie/pkg/config"
	"github.com/bjerkio/stackie/pkg/stackie"
	"github.com/bradleyjkemp/cupaloy/v2"
	"github.com/spf13/afero"
)

var UserConfig = `
pulumi:
    - name: not-a-real-profile
	  accessToken: not-a-real-token
`

func prepareUserConfig(c config.Config, fs afero.Fs) error {
	p := path.Join("project", ".stackie.local.yml")
	err := afero.WriteFile(fs, p, []byte(UserConfig), 0644)
	if err != nil {
		return err
	}
	return nil
}

var ProjectConfig = `
name: Stackie Test
stacks:
    - name: dev
      pulumi:
        stackName: dev
        cloudUrl: gs://hello-world
    - name: dev
      pulumi:
        stackName: dev
        cloudUrl: gs://hello-world
`

func prepareProjectConfig(c config.Config, fs afero.Fs) error {
	p := path.Join("project", ".stackie.yml")
	err := afero.WriteFile(fs, p, []byte(ProjectConfig), 0644)
	if err != nil {
		return err
	}
	return nil
}

var UserProjectConfig = `
activeStack: dev
`

func prepareUserProjectConfig(c config.Config, fs afero.Fs, t *testing.T) error {
	homeDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	userProjectPath := path.Join(homeDir, stackie.CONFIG_FOLDER, stackie.CONFIG_NAME)
	fs.MkdirAll(path.Join(homeDir, stackie.CONFIG_FOLDER), 0755)

	err = afero.WriteFile(fs, userProjectPath, []byte(UserProjectConfig), 0644)
	if err != nil {
		return err
	}

	if !c.Exists(userProjectPath) {
		t.Errorf("Expected the .stackie.yml config to exist")
	}

	return nil
}

func TestStackieRead(t *testing.T) {
	c := config.New()
	fs := afero.NewMemMapFs()
	c.SetFS(fs)

	fs.MkdirAll("project", 0755)

	err := prepareUserConfig(c, fs)
	if err != nil {
		t.Error(err)
	}

	err = prepareProjectConfig(c, fs)
	if err != nil {
		t.Error(err)
	}

	err = prepareUserProjectConfig(c, fs, t)
	if err != nil {
		t.Error(err)
	}

	conf, err := stackie.ReadConfigs(c)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(conf.UserProjectConfig)

	cupaloy.SnapshotT(t, conf)
}
