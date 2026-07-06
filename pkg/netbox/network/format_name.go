package network

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Interface) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Name)
	}

	return i.Name
}
