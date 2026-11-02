package module_bay

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (b *Bay) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", b.Name)
	}

	return b.Name
}
