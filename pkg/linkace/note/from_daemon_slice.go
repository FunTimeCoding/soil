package note

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemonSlice(v []client.Note) []*Note {
	var result []*Note

	for _, e := range v {
		result = append(result, FromDaemon(e))
	}

	return result
}
