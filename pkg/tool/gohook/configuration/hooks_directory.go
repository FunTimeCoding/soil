package configuration

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"path/filepath"
)

func (c *Configuration) HooksDirectory() string {
	return filepath.Join(c.Root, constant.Directory, constant.HooksDirectory)
}
