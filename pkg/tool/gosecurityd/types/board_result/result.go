package board_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

type Result struct {
	Total          int
	Supported      int
	Tracked        int
	Untracked      int
	Stale          int
	Upgradable     int
	Critical       int
	WithFix        int
	NoFix          int
	Kernels        map[string]int
	Rows           []*system_result.Summary
	CoverageCounts map[string]int
	ScoreCounts    map[float64]int
	LabelCounts    map[string]map[string]int
}
