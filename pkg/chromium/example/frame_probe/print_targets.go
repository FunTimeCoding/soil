package frame_probe

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/notation"
)

func printTargets(c *chromium.Client) {
	console.Line("=== Target.getTargets ===")
	targets, e := c.Targets()
	errors.PanicOnError(e)

	for _, t := range targets {
		console.Line(notation.MarshalIndent(t))
	}
}
