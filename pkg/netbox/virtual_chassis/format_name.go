package virtual_chassis

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Chassis) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", c.Name)
	}

	return c.Name
}
