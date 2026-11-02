package mock_client

import "github.com/funtimecoding/soil/pkg/docker/hub/tag"

func (c *Client) SeedTags(
	image string,
	names ...string,
) {
	for _, name := range names {
		c.tags[image] = append(c.tags[image], tag.NewNamed(name))
	}

	if _, okay := c.tags[image]; !okay {
		c.tags[image] = nil
	}
}
