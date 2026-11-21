package board_result

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"

func New() *Result {
	return &Result{
		Rows:           []*system_result.Summary{},
		Kernels:        map[string]int{},
		CoverageCounts: map[string]int{},
		ScoreCounts:    map[float64]int{},
		LabelCounts:    map[string]map[string]int{},
	}
}
