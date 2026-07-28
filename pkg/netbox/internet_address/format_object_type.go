package internet_address

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/netbox/constant"
)

func (a *Address) formatObjectType(f *option.Format) string {
	result := a.ObjectType

	if result == "" {
		result = constant.NoObjectType

		if f.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
