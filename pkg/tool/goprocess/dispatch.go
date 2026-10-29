package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func Dispatch(
	c face.Client,
	command string,
	arguments []string,
) error {
	x := context.Background()

	switch command {
	case "start":
		return each(arguments, func(n string) error { return c.Start(x, n) })
	case "stop":
		return each(arguments, func(n string) error { return c.Stop(x, n) })
	case "restart":
		return each(arguments, func(n string) error { return c.Restart(x, n) })
	case "restart-all":
		return restartAll(x, c)
	case "reload-procfile":
		return c.ReloadProcfile(x)
	case "reload-environment":
		return c.ReloadEnvironment(x)
	case "log":
		return log(x, c, arguments)
	case "list":
		return list(x, c)
	case "status":
		return status(x, c)
	default:
		return validation.New("unknown command %s", command)
	}
}
