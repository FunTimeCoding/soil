package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/note"

func (c *Client) NotesByLink(_ int) ([]*note.Note, error) {
	return nil, c.Failure
}
