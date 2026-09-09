package mock_client

import "github.com/funtimecoding/soil/pkg/opnsense/log_entry"

func (c *Client) Log(_ int) ([]*log_entry.Entry, error) {
	return nil, nil
}
