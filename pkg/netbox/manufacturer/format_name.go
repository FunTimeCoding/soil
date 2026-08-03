package manufacturer

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Manufacturer) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", m.Name)
	}

	return m.Name
}
