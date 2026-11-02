package physical_address

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Address) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", a.Name)
	}

	return a.Name
}
