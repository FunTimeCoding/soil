package physical_address

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Address) formatObjectType(f *option.Format) string {
	result := a.ObjectType

	if result == "" {
		result = NoObjectType

		if f.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
