package list

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (l *List) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", l.Name)
	}

	return l.Name
}
