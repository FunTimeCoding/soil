package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"slices"
)

func (a *Alert) formatSeverity(f *option.Format) string {
	result := a.Severity

	if f.UseColor {
		if slices.Contains(constant.RedSeverities, result) {
			result = console.Red("%s", result)
		} else if slices.Contains(constant.YellowSeverities, result) {
			result = console.Yellow("%s", result)
		}
	}

	if result == constant.NoneSeverity {
		result = "no severity"
	}

	return result
}
