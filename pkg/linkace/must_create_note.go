package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/note"
)

func (c *Client) MustCreateNote(
	linkIdentifier int,
	text string,
) *note.Note {
	result, e := c.CreateNote(linkIdentifier, text)
	errors.PanicOnError(e)

	return result
}
