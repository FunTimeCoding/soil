package alertmanager

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustCreate(
	name string,
	instance string,
	summary string,
	description string,
	generator string,
) {
	errors.PanicOnError(
		c.Create(name, instance, summary, description, generator),
	)
}
