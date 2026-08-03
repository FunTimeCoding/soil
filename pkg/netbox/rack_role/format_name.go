package rack_role

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Role) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", r.Name)
	}

	return r.Name
}
