package client

import (
	"github.com/funtimecoding/soil/pkg/directory"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
)

func add(
	c *directory.Client,
	name string,
	attributes map[string][]string,
) {
	e := c.Add(name, attributes)

	if conflict.Is(e) {
		return
	}

	errors.PanicOnError(e)
}
