package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func list(
	x context.Context,
	c face.Client,
	t *terminal.Terminal,
) {
	v, e := c.ListProcesses(x)

	if e != nil {
		t.EmitDone(e)

		return
	}

	names := make([]string, len(v))

	for i, p := range v {
		names[i] = p.Name
	}

	t.EmitValue(names, nil)
}
