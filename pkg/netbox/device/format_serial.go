package device

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (d *Device) formatSerial(f *option.Format) string {
	if d.Serial == "" {
		if f.UseColor {
			return console.Red("%s", constant.NoSerial)
		}

		return constant.NoSerial
	}

	return d.Serial
}
