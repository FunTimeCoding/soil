package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) MustNotesByLinkPage(
	linkIdentifier int,
	p int,
) *page.Page[*note.Note] {
	result, e := c.NotesByLinkPage(linkIdentifier, p)
	errors.PanicOnError(e)

	return result
}
