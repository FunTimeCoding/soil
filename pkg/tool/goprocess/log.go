package goprocess

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func log(
	x context.Context,
	c face.Client,
	arguments []string,
) error {
	if len(arguments) == 0 {
		return validation.New("log requires a process name")
	}

	name := arguments[0]
	option := ""

	if len(arguments) > 1 {
		option = arguments[1]
	}

	switch option {
	case "clear":
		return c.ClearLog(x, name)
	case "", "all":
		v, e := c.Log(x, name, option == "all")

		if e != nil {
			return e
		}

		for _, l := range v.Lines {
			console.Line(l)
		}

		if v.Older > 0 {
			console.Format("(%d older lines, use: log %s all)\n", v.Older, name)
		}

		return nil
	default:
		return validation.New("unknown log option %s", option)
	}
}
