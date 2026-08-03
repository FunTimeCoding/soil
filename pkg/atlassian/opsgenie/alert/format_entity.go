package alert

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatEntity(f *option.Format) string {
	result := a.Entity

	if result != "" && f.UseColor {
		result = constant.Cyan("%s", result)
	}

	return result
}
