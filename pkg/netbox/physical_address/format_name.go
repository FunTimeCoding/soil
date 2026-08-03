package physical_address

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Address) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", a.Name)
	}

	return a.Name
}
