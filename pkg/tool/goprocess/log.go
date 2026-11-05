package goprocess

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/face"
)

func log(
	x context.Context,
	c face.Client,
	t *terminal.Terminal,
	arguments []string,
) {
	if len(arguments) == 0 {
		t.EmitDone(validation.New("log requires a process name"))

		return
	}

	name := arguments[0]
	option := ""

	if len(arguments) > 1 {
		option = arguments[1]
	}

	switch option {
	case "clear":
		t.EmitDone(c.ClearLog(x, name))
	case "", "all":
		v, e := c.GetLog(x, name, option == "all")

		if e != nil {
			t.EmitDone(e)

			return
		}

		for _, l := range v.Lines {
			t.Line(l)
		}

		if v.Older > 0 {
			t.Noticeln(
				fmt.Sprintf("(%d older lines, use: log %s all)", v.Older, name),
			)
		}
	default:
		t.EmitDone(validation.New("unknown log option %s", option))
	}
}
