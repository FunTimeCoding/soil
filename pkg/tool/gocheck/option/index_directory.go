package option

import "github.com/funtimecoding/soil/pkg/source/index"

func (c *Check) IndexDirectory() string {
	if c.Index == "" {
		return index.DefaultDirectory()
	}

	return c.Index
}
