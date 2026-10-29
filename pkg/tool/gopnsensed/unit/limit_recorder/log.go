package limit_recorder

import "github.com/funtimecoding/soil/pkg/opnsense/log_entry"

func (c *Client) Log(limit int) ([]*log_entry.Entry, error) {
	c.limit = limit

	return nil, nil
}
