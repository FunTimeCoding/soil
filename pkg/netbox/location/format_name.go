package location

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (l *Location) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", l.Name)
	}

	return l.Name
}
