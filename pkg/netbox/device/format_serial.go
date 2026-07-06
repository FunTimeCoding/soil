package device

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (d *Device) formatSerial(f *option.Format) string {
	if d.Serial == "" {
		if f.UseColor {
			return console.Red("%s", NoSerial)
		}

		return NoSerial
	}

	return d.Serial
}
