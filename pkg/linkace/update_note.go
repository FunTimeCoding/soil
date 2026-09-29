package linkace

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) UpdateNote(
	identifier int,
	text string,
) (*note.Note, error) {
	var r response.Note

	if e := c.basic.Patch(
		fmt.Sprintf("notes/%d", identifier),
		map[string]any{"note": text},
		&r,
	); e != nil {
		return nil, e
	}

	return note.New(r), nil
}
