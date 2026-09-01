package frame_probe

import (
	"github.com/funtimecoding/soil/pkg/chromium"
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/notation"
)

func printTargets(c *chromium.Client) {
	console.Line("=== Target.getTargets ===")

	for _, t := range c.Targets() {
		console.Line(notation.MarshalIndent(t))
	}
}
