package log_entry

import "github.com/funtimecoding/soil/pkg/opnsense/response"

func NewSlice(v []response.LogEntry) []*Entry {
	var result []*Entry

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
