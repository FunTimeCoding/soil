package model_context

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
)

func retargetLines(reports []*retarget_report.Report) []string {
	var result []string

	for _, r := range reports {
		result = append(result, fmt.Sprintf("#%d %s", r.Identifier, r.Name))

		for _, c := range r.Changes {
			switch {
			case c.Reason != "":
				result = append(
					result,
					fmt.Sprintf(
						"  line %d: %s - %s",
						c.Line,
						c.Before,
						c.Reason,
					),
				)
			case c.Line == 0:
				result = append(
					result,
					fmt.Sprintf("  base: %s → %s", c.Before, c.After),
				)
			default:
				result = append(
					result,
					fmt.Sprintf(
						"  line %d: %s → %s",
						c.Line,
						c.Before,
						c.After,
					),
				)
			}
		}
	}

	return result
}
