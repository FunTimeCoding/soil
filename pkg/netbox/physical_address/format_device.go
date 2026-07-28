package physical_address

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (a *Address) formatDevice(f *option.Format) string {
	if a.Interface == nil {
		return ""
	}

	result := a.Interface.Device.GetName()

	if result == "" {
		result = constant.NoDevice

		if f.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
