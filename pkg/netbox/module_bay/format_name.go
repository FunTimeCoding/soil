package module_bay

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (b *Bay) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", b.Name)
	}

	return b.Name
}
