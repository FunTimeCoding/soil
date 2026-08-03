package alert

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"slices"
)

func (a *Alert) formatSeverity(f *option.Format) string {
	result := a.Severity

	if f.UseColor {
		if slices.Contains(constant.RedSeverities, result) {
			result = consoleConstant.Red("%s", result)
		} else if slices.Contains(constant.YellowSeverities, result) {
			result = consoleConstant.Yellow("%s", result)
		}
	}

	if result == constant.NoneSeverity {
		result = "no severity"
	}

	return result
}
