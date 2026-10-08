package search_index

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func kindsOrDefault(kinds []string) []string {
	if len(kinds) == 0 {
		return []string{constant.BlockMessage}
	}

	return kinds
}
