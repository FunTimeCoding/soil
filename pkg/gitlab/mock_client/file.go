package mock_client

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) File(
	_ int64,
	_ string,
	name string,
) (*gitlab.File, error) {
	f, exists := c.files[name]

	if !exists {
		return nil, fmt.Errorf("file: %s: %w", name, constant.ErrorNotFound)
	}

	return f, nil
}
