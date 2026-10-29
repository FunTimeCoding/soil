package mock_client

import "github.com/funtimecoding/soil/pkg/nextcloud/usage"

func (c *Client) Fetch() (*usage.Usage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return usage.New(c.usage.Files, c.usage.Shares), nil
}
