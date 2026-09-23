package fixture

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func EntriesByKind(
	entries []queue.Entry,
	kind string,
) []queue.Entry {
	var result []queue.Entry

	for _, e := range entries {
		if e.Kind == kind {
			result = append(result, e)
		}
	}

	return result
}
