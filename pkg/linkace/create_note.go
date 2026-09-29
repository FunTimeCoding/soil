package linkace

import (
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/response"
)

func (c *Client) CreateNote(
	linkIdentifier int,
	text string,
) (*note.Note, error) {
	var r response.Note

	if e := c.basic.Post(
		"notes",
		map[string]any{
			"link_id":    linkIdentifier,
			"note":       text,
			"visibility": 3,
		},
		&r,
	); e != nil {
		return nil, e
	}

	return note.New(r), nil
}
