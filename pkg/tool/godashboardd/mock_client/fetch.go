package mock_client

import "github.com/funtimecoding/soil/pkg/nextcloud/usage"

func (c *Client) Fetch() (*usage.Usage, error) {
	return nil, nil
}
