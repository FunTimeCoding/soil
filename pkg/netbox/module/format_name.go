package module

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (m *Module) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", m.Name)
	}

	return m.Name
}
