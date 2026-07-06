package power_feed

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (f *Feed) formatName(o *option.Format) string {
	if o.UseColor {
		return console.Cyan("%s", f.Name)
	}

	return f.Name
}
