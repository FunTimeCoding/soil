package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatStatus(f *option.Format) string {
	var result string

	if f.UseColor {
		if a.Status != ClosedStatus {
			result = console.Red("%s", a.Status)
		} else {
			result = a.Status
		}
	} else {
		result = a.Status
	}

	return result
}
