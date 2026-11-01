package mock_client

import "github.com/google/go-github/v93/github"

func (c *Client) Notifications() ([]*github.Notification, error) {
	return nil, c.Failure
}
