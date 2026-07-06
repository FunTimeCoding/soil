package channel

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Channel) Format(f *option.Format) string {
	return status.New(f).Integer64(c.Identifier).String(c.Name).Format()
}
