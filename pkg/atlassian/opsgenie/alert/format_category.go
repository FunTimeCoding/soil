package alert

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatCategory(f *option.Format) string {
	result := a.Category

	if result != "" && f.UseColor {
		result = constant.Cyan("%s", result)
	}

	return result
}
