package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/note"
)

func (c *Client) MustNotesByLink(linkIdentifier int) []*note.Note {
	result, e := c.NotesByLink(linkIdentifier)
	errors.PanicOnError(e)

	return result
}
