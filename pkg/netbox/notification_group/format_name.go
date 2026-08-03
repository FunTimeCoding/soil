package notification_group

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (g *Group) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", g.Name)
	}

	return g.Name
}
