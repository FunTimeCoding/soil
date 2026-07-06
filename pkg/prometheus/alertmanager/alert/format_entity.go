package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatEntity(f *option.Format) string {
	if a.Entity == "" {
		return ""
	}

	if f.UseColor {
		return console.Cyan("%s", a.Entity)
	}

	return a.Entity
}
