package dependency

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (d *Dependency) formatPath(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", d.Path)
	}

	return d.Path
}
