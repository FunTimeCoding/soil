package delivery

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/queue"

func group(entries []queue.Entry) map[string][]queue.Entry {
	result := map[string][]queue.Entry{}

	for _, e := range entries {
		result[e.Kind] = append(result[e.Kind], e)
	}

	return result
}
