package mock_client

import (
	"github.com/funtimecoding/soil/pkg/linkace/note"
	"github.com/funtimecoding/soil/pkg/linkace/page"
)

func (c *Client) NotesByLinkPage(
	_ int,
	_ int,
) (*page.Page[*note.Note], error) {
	return &page.Page[*note.Note]{LastPage: 1}, nil
}
