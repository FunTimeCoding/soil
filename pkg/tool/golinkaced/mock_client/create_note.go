package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/note"

func (c *Client) CreateNote(
	_ int,
	_ string,
) (*note.Note, error) {
	return nil, nil
}
