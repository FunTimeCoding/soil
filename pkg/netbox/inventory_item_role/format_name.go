package inventory_item_role

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Role) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", r.Name)
	}

	return r.Name
}
