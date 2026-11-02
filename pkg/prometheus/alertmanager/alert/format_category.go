package alert

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatCategory(f *option.Format) string {
	if a.Category == "" {
		return ""
	}

	if f.UseColor {
		return consoleConstant.Cyan("%s", a.Category)
	}

	return a.Category
}
