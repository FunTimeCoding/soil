package custom_field

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (f *Field) formatName(o *option.Format) string {
	if o.UseColor {
		return console.Cyan("%s", f.Name)
	}

	return f.Name
}
