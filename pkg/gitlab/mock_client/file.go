package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/file"
)

func (c *Client) File(
	_ int64,
	_ string,
	name string,
) (*file.File, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	f, exists := c.files[name]

	if !exists {
		return nil, not_found.New("file", name)
	}

	return f, nil
}
