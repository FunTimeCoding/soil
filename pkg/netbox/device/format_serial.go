package device

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (d *Device) formatSerial(f *option.Format) string {
	if d.Serial == "" {
		if f.UseColor {
			return consoleConstant.Red("%s", constant.NoSerial)
		}

		return constant.NoSerial
	}

	return d.Serial
}
