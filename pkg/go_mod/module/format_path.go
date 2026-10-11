package module

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Module) formatPath(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", m.Path)
	}

	return m.Path
}
