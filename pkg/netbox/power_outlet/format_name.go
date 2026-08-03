package power_outlet

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (o *Outlet) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", o.Name)
	}

	return o.Name
}
