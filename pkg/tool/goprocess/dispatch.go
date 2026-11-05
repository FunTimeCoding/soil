package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func Dispatch(
	c face.Client,
	t *terminal.Terminal,
	command string,
	arguments []string,
) {
	x := context.Background()

	switch command {
	case "start":
		t.EmitDone(
			each(
				arguments,
				func(n string) error { return c.StartProcess(x, n) },
			),
		)
	case "stop":
		t.EmitDone(
			each(
				arguments,
				func(n string) error { return c.StopProcess(x, n) },
			),
		)
	case "restart":
		t.EmitDone(
			each(
				arguments,
				func(n string) error { return c.RestartProcess(x, n) },
			),
		)
	case "restart-all":
		restartAll(x, c, t)
	case "reload-procfile":
		t.EmitDone(c.ReloadProcfile(x))
	case "reload-environment":
		t.EmitDone(c.ReloadEnvironment(x))
	case "log":
		log(x, c, t, arguments)
	case "list":
		list(x, c, t)
	case "status":
		t.EmitValue(c.ListProcesses(x))
	default:
		t.EmitDone(validation.New("unknown command %s", command))
	}
}
