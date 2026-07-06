package tag

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Tag) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", t.Name)
	}

	return t.Name
}
