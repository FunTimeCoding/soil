package physical_address

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Address) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", a.Name)
	}

	return a.Name
}
