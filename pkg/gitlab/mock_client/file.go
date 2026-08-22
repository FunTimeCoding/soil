package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) File(
	_ int64,
	_ string,
	name string,
) (*gitlab.File, error) {
	f, exists := c.files[name]

	if !exists {
		return nil, not_found.New("file", name)
	}

	return f, nil
}
