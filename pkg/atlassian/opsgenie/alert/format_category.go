package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatCategory(f *option.Format) string {
	result := a.Category

	if result != "" && f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
