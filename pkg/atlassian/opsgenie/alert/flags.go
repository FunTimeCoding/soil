package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) flags(f *option.Format) []string {
	var result []string

	if !f.HasTag(consoleConstant.TagDense) {
		if !a.Acknowledged {
			flag := constant.OpsgenieUnacknowledgedFlag

			if f.UseColor {
				flag = consoleConstant.Red("%s", flag)
			}

			result = append(result, flag)
		}

		if !a.Seen {
			flag := constant.OpsgenieUnseenFlag

			if f.UseColor {
				flag = consoleConstant.Red("%s", flag)
			}

			result = append(result, flag)
		}
	}

	if a.Snoozed {
		flag := constant.OpsgenieSnoozedFlag

		if f.UseColor {
			flag = consoleConstant.Red("%s", flag)
		}

		result = append(result, flag)
	}

	return result
}
