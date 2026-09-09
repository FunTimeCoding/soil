package discussion

import "github.com/funtimecoding/soil/pkg/gitlab/note"

type Discussion struct {
	Identifier string
	Individual bool
	Notes      []*note.Note
}
