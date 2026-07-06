package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (a *Alert) flags(f *option.Format) []string {
	var result []string

	if !f.HasTag(tag.Dense) {
		if !a.Acknowledged {
			flag := UnacknowledgedFlag

			if f.UseColor {
				flag = console.Red("%s", flag)
			}

			result = append(result, flag)
		}

		if !a.Seen {
			flag := UnseenFlag

			if f.UseColor {
				flag = console.Red("%s", flag)
			}

			result = append(result, flag)
		}
	}

	if a.Snoozed {
		flag := SnoozedFlag

		if f.UseColor {
			flag = console.Red("%s", flag)
		}

		result = append(result, flag)
	}

	return result
}
