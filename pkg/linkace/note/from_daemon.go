package note

import "github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"

func FromDaemon(v client.Note) *Note {
	return &Note{Identifier: int(v.Identifier), Text: v.Text}
}
