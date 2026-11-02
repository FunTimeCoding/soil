package dependency

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (d *Dependency) formatPath(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", d.Path)
	}

	return d.Path
}
