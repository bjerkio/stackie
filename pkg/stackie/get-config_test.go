package stackie

import (
	"testing"

	"github.com/spf13/afero"
)

func TestGetConfigPath(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.MkdirAll("project", 0755)
	afero.WriteFile(fs, "project/.stackie.json", []byte(""), 0644)

	c := &ConfigFS{
		fs,
	}

	p, err := c.getPath("project")
	if err != nil {
		t.Error(err)
	}

	if p != "project/.stackie.json" {
		t.Errorf("getConfigPath failed, expected %v, got %v", ".stackie.json", p)
	}
}
