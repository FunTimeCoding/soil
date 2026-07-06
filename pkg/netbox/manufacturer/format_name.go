package manufacturer

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Manufacturer) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", m.Name)
	}

	return m.Name
}
