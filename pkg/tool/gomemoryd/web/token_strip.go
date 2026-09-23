package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/token_summary"
)

func tokenStrip(
	m *store.Memory,
	statistic *token_summary.Statistic,
	summary *token_summary.Summary,
) []string {
	if statistic == nil {
		return nil
	}

	return []string{
		fmt.Sprintf("block %d", statistic.Block),
		fmt.Sprintf("description %d", statistic.Description),
		fmt.Sprintf(
			"heavier than %d%% of memories",
			token_summary.BlockRank(summary, statistic.Block),
		),
		reachText(m),
	}
}
