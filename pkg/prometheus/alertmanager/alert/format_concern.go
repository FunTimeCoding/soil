package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (a *Alert) formatConcern(f *option.Format) string {
	if len(a.concern) == 0 {
		return ""
	}

	result := join.Comma(a.concern)

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
