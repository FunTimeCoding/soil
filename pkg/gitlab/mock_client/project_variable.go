package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (c *Client) ProjectVariable(
	_ int64,
	_ string,
) (*variable.Variable, error) {
	return nil, nil
}
