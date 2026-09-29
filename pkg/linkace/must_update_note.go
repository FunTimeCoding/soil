package linkace

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/linkace/note"
)

func (c *Client) MustUpdateNote(
	identifier int,
	text string,
) *note.Note {
	result, e := c.UpdateNote(identifier, text)
	errors.PanicOnError(e)

	return result
}
