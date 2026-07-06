package cluster

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Cluster) Format(f *option.Format) string {
	return status.New(f).String(c.formatName(f)).RawList(c.Raw).Format()
}
