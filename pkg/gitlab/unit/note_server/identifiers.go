package note_server

import "github.com/funtimecoding/soil/pkg/gitlab/note"

func Identifiers(v []*note.Note) []int64 {
	var result []int64

	for _, n := range v {
		result = append(result, n.Identifier)
	}

	return result
}
