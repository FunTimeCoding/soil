package token_summary

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func BlockRank(
	summary *Summary,
	block int,
) int {
	if len(summary.Statistic) == 0 {
		return 0
	}

	lighter := 0

	for _, t := range summary.Statistic {
		if t.Block < block {
			lighter++
		}
	}

	return lighter * constant.WholePercent / len(summary.Statistic)
}
