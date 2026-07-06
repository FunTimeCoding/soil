package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatName(f *option.Format) string {
	result := a.Name

	if f.UseColor {
		// Not meant to be cyan, because entity and category are cyan
		result = console.Yellow("%s", result)
	}

	return result
}
