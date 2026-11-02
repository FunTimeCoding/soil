package alert

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatEntity(f *option.Format) string {
	if a.Entity == "" {
		return ""
	}

	if f.UseColor {
		return consoleConstant.Cyan("%s", a.Entity)
	}

	return a.Entity
}
