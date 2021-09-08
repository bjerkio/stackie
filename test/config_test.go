package test

import (
	"path"
	"testing"

	"github.com/bjerkio/stackie/pkg/config"
	"github.com/spf13/afero"
)

func TestGetConfigPath(t *testing.T) {
	c := config.New()

	var testConfigNames = func(configNames []config.FileName) {
		for _, name := range configNames {
			fs := afero.NewMemMapFs()
			c.SetFS(fs)
			fs.MkdirAll("project", 0755)

			sf := path.Join("project", name.FileName)

			afero.WriteFile(fs, sf, []byte(""), 0644)

			p, _, err := c.GetPath("project", configNames)
			if err != nil {
				t.Error(err)
			}

			if p != sf {
				t.Errorf("GetPath failed, expected %v, got %v", sf, p)
			}

			t.Logf("GetPath found %v", name.FileName)
		}
	}

	userConfig := config.GetUserFileNames()
	testConfigNames(userConfig)

	projectConfig := config.GetProjectFileNames()
	testConfigNames(projectConfig)

	fs := afero.NewMemMapFs()
	fs.MkdirAll("second-project", 0755)

	_, _, err := c.GetPath("second-project", userConfig)
	if err == nil {
		t.Errorf("getConfigPath failed, expected second-project to fail")
	}
}
