package mock_outpost_source

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"

func (c *Client) Add(
	name string,
	origin string,
	source string,
) {
	c.services = append(
		c.services,
		client.Service{
			Name:       name,
			State:      "running",
			Origin:     origin,
			Source:     &source,
			Deliberate: true,
		},
	)
}
