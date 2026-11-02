package wireless_network_group

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (g *Group) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", g.Name)
	}

	return g.Name
}
