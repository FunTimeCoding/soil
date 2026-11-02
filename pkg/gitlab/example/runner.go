package example

import (
	"context"
	"github.com/funtimecoding/soil/pkg/console"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Runner() {
	g := gitlab.NewEnvironment()
	defer g.Close()
	x := context.Background()
	f := consoleConstant.ExtendedColorFormat.Copy().Raw()

	for _, r := range g.MustRunners(x, true) {
		r = g.MustRunner(x, r.Identifier)
		console.Line(r.Format(f))
	}
}
