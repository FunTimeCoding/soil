package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func status(
	x context.Context,
	c face.Client,
) error {
	v, e := c.Processes(x)

	if e != nil {
		return e
	}

	for _, p := range v {
		prefix := " "

		if p.Running {
			prefix = "*"
		}

		console.Format("%s%s\n", prefix, p.Name)
	}

	return nil
}
