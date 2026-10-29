package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"strings"
)

func alertLines(alerts []client.AlertsResponse) string {
	var result []string

	for _, a := range alerts {
		end := ""

		if a.End != nil {
			end = *a.End
		}

		result = append(
			result,
			fmt.Sprintf(
				"%d %s %s %s %s",
				a.Identifier,
				a.Name,
				a.Status,
				a.Start,
				end,
			),
		)
	}

	return strings.Join(result, "\n")
}
