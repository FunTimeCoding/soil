package gomemory

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func statisticTable(s *client.TokenSummary) string {
	line := make([]string, 0, len(s.Statistic)+8)
	line = append(line, fmt.Sprintf("%-6s %-6s  %s", "block", "descr", "name"))

	for _, t := range s.Statistic {
		line = append(
			line,
			fmt.Sprintf(
				"%-6d %-6d  %s%s",
				t.Block,
				t.Description,
				t.Name,
				tagSuffix(t.Tags),
			),
		)
	}

	line = append(
		line,
		"",
		fmt.Sprintf(
			"%-12s %7s %9s %12s %8s",
			"",
			"median",
			"ninetieth",
			"ninetyninth",
			"maximum",
		),
		spreadRow("block", s.Block),
		spreadRow("description", s.Description),
		"",
		fmt.Sprintf(
			"%d memories, %d tokens if every block were loaded",
			len(s.Statistic)+s.Withheld,
			s.Block.Total,
		),
		withheldNote(s.Withheld),
	)

	return join.NewLine(line)
}
