package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocess/result"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func restartAll(
	x context.Context,
	c face.Client,
	t *terminal.Terminal,
) {
	count, e := c.RestartAll(x)

	if e != nil {
		t.EmitDone(e)

		return
	}

	t.EmitValue(result.NewRestart(count), nil)
}
