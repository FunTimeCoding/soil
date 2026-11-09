package unit

import "github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"

func clientKindEntries(
	entries []client.QueueEntry,
	kind string,
) []client.QueueEntry {
	var result []client.QueueEntry

	for _, e := range entries {
		if e.Kind == kind {
			result = append(result, e)
		}
	}

	return result
}
