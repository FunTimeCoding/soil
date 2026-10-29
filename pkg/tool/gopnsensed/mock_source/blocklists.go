package mock_source

import "github.com/funtimecoding/soil/pkg/opnsense/blocklist"

func (c *Client) Blocklists(_ string) ([]*blocklist.Blocklist, error) {
	return nil, nil
}
