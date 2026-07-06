package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatCategory(f *option.Format) string {
	if a.Category == "" {
		return ""
	}

	if f.UseColor {
		return console.Cyan("%s", a.Category)
	}

	return a.Category
}
