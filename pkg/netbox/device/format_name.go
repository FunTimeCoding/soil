package device

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (d *Device) formatName(f *option.Format) string {
	if d.Name == "" {
		if f.UseColor {
			return consoleConstant.Yellow("%s", constant.NoName)
		}

		return constant.NoName
	}

	if f.UseColor {
		return consoleConstant.Cyan("%s", d.Name)
	}

	return d.Name
}
