package module

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Module) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", m.Name)
	}

	return m.Name
}
