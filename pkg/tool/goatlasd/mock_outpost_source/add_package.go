package mock_outpost_source

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/generated/client"

func (c *Client) AddPackage(
	name string,
	origin string,
	source string,
	packageName string,
	version string,
) {
	c.services = append(
		c.services,
		client.Service{
			Name:       name,
			State:      "running",
			Origin:     origin,
			Source:     &source,
			Package:    &packageName,
			Version:    &version,
			Deliberate: true,
		},
	)
}
