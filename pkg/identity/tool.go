package identity

import (
	"github.com/funtimecoding/soil/pkg/identity/paragraph"
	"github.com/funtimecoding/soil/pkg/stamp"
)

type Tool struct {
	name         string
	description  string
	usage        string
	instructions string
	stamp        *stamp.Stamp
	paragraphs   []*paragraph.Paragraph
}
