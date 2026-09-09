package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (c *Client) CreateProjectVariable(
	_ int64,
	_ string,
	_ string,
	_ bool,
	_ bool,
	_ bool,
) (*variable.Variable, error) {
	return nil, nil
}
