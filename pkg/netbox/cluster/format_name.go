package cluster

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Cluster) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", c.Name)
	}

	return c.Name
}
