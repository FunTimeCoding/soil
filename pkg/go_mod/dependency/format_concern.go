package dependency

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (d *Dependency) formatConcern(f *option.Format) string {
	if len(d.concern) == 0 {
		return ""
	}

	result := join.Comma(d.concern)

	if f.UseColor {
		result = constant.Yellow("%s", result)
	}

	return result
}
