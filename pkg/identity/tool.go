package identity

import "github.com/funtimecoding/soil/pkg/identity/paragraph"

type Tool struct {
	name         string
	description  string
	usage        string
	instructions string
	paragraphs   []*paragraph.Paragraph
}
