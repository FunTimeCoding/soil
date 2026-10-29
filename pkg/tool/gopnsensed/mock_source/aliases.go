package mock_source

import "github.com/funtimecoding/soil/pkg/opnsense/alias"

func (c *Client) Aliases(_ string) ([]*alias.Alias, error) {
	return nil, nil
}
