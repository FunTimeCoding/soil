package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/alias"

func (c *Client) Aliases(_ string) ([]*alias.Alias, error) {
	return nil, nil
}
