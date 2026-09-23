package token_summary

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"

func at(
	sorted []int,
	percentile int,
) int {
	return sorted[(len(sorted)-1)*percentile/constant.WholePercent]
}
