package alert

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) flags(f *option.Format) []string {
	var result []string

	if !f.HasTag(constant.TagDense) {
		if !a.Acknowledged {
			flag := atlassian.OpsgenieUnacknowledgedFlag

			if f.UseColor {
				flag = constant.Red("%s", flag)
			}

			result = append(result, flag)
		}

		if !a.Seen {
			flag := atlassian.OpsgenieUnseenFlag

			if f.UseColor {
				flag = constant.Red("%s", flag)
			}

			result = append(result, flag)
		}
	}

	if a.Snoozed {
		flag := atlassian.OpsgenieSnoozedFlag

		if f.UseColor {
			flag = constant.Red("%s", flag)
		}

		result = append(result, flag)
	}

	return result
}
