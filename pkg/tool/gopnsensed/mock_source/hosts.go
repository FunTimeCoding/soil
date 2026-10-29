package mock_source

import "github.com/funtimecoding/soil/pkg/opnsense/host"

func (c *Client) Hosts(_ string) ([]*host.Host, error) {
	return nil, nil
}
