package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/host"

func (c *Client) Hosts(_ string) ([]*host.Host, error) {
	return nil, nil
}
