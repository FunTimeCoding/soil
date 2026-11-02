package virtual_device_context

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Context) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", c.Name)
	}

	return c.Name
}
