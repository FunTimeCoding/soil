package physical_address

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
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
			result = consoleConstant.Red("%s", result)
		}
	}

	return result
}
