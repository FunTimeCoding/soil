package container

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Container) Format(f *option.Format) string {
	s := status.New(f).String(
		c.formatName(f),
		c.Repository,
	).RawList(c.Raw)

	return s.Format()
}
