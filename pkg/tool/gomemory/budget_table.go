package gomemory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func budgetTable(
	d *client.ProfileDetail,
	r *client.ProfileResponse,
) string {
	return join.NewLine(
		[]string{
			"",
			fmt.Sprintf(
				"%-13s %7s %7s %8s",
				"tier",
				"tokens",
				"shown",
				"trimmed",
			),
			budgetRow("always", d.AlwaysTokens, len(r.Always), 0),
			budgetRow("index", d.IndexTokens, len(r.Index), d.IndexTrimmed),
			budgetRow(
				"completions",
				d.CompletionTokens,
				countCompletions(r),
				d.CompletionsTrimmed,
			),
			budgetRow(
				"impressions",
				d.ImpressionTokens,
				countImpressions(r),
				d.ImpressionsTrimmed,
			),
			budgetRow(
				"relevant",
				d.RelevantTokens,
				countRelevant(r),
				d.RelevantTrimmed,
			),
			"",
			fmt.Sprintf("%-13s %7d of %d", "total", d.TotalTokens, d.Budget),
			hiddenNote(d.Hidden),
		},
	)
}
