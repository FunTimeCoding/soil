package device

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (d *Device) formatName(f *option.Format) string {
	if d.Name == "" {
		if f.UseColor {
			return console.Yellow("%s", constant.NoName)
		}

		return constant.NoName
	}

	if f.UseColor {
		return console.Cyan("%s", d.Name)
	}

	return d.Name
}
