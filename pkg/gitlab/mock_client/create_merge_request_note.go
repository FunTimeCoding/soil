package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/note"

func (c *Client) CreateMergeRequestNote(
	_ int64,
	_ int64,
	_ string,
) (*note.Note, error) {
	return nil, nil
}
