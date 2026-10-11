package mock_client

import "github.com/funtimecoding/soil/pkg/strings/join"

func (c *Client) SeedVersions(
	repository string,
	name string,
	versions ...string,
) {
	c.versions[join.Slash([]string{repository, name})] = versions
}
