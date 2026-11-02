package link

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (l *Link) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", l.Title)
	}

	return l.Title
}
