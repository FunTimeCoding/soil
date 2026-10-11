package mock_client

import "github.com/funtimecoding/soil/pkg/strings/join"

func (c *Client) Versions(
	repository string,
	name string,
) ([]string, error) {
	if c.fail != nil {
		return nil, c.fail
	}

	return c.versions[join.Slash([]string{repository, name})], nil
}
