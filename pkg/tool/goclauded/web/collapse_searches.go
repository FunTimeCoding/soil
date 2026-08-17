package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
)

func collapseSearches(loads []context_load.Load) []context_load.Load {
	var result []context_load.Load
	counts := map[string]int{}

	for _, entry := range loads {
		if entry.Kind != constant.LoadKindSearch {
			result = append(result, entry)

			continue
		}

		counts[entry.CallIdentifier]++

		if counts[entry.CallIdentifier] > 1 {
			continue
		}

		entry.Reference = join.Empty(`"`, entry.Query, `"`)
		result = append(result, entry)
	}

	for i := range result {
		if result[i].Kind != constant.LoadKindSearch {
			continue
		}

		result[i].Name = fmt.Sprintf(
			"%d hits",
			counts[result[i].CallIdentifier],
		)
	}

	return result
}
