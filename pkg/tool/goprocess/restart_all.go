package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func restartAll(
	x context.Context,
	c face.Client,
) error {
	count, e := c.RestartAll(x)

	if e != nil {
		return e
	}

	console.Format("restarting %d processes in the background\n", count)

	return nil
}
