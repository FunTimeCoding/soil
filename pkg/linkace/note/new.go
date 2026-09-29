package note

import "github.com/funtimecoding/soil/pkg/linkace/response"

func New(v response.Note) *Note {
	return &Note{
		Identifier:     v.Identifier,
		LinkIdentifier: v.LinkIdentifier,
		Text:           v.Text,
		Visibility:     v.Visibility,
	}
}
