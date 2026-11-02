package configuration

import (
	gitConstant "github.com/funtimecoding/soil/pkg/git/constant"
	"path/filepath"
)

func (c *Configuration) HooksDirectory() string {
	return filepath.Join(
		c.Root,
		gitConstant.Directory,
		gitConstant.HooksDirectory,
	)
}
