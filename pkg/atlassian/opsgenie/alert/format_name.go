package alert

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatName(f *option.Format) string {
	result := a.Name

	if f.UseColor {
		result = constant.Yellow("%s", result)
	}

	return result
}
