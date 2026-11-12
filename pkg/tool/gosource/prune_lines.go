package gosource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosource/constant"
)

func PruneLines(removed []string) []string {
	if len(removed) == 0 {
		return []string{constant.NothingToPrune}
	}

	result := make([]string, 0, len(removed))

	for _, name := range removed {
		result = append(result, fmt.Sprintf(constant.RemovedFormat, name))
	}

	return result
}
