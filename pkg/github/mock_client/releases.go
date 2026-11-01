package mock_client

import "github.com/funtimecoding/soil/pkg/github/release"

func (c *Client) Releases(_ string, _ string) ([]*release.Release, error) {
	return nil, c.Failure
}
