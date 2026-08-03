package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/gitlab"
)

func Runner() {
	g := gitlab.NewEnvironment()
	f := constant.ExtendedColorFormat.Copy().Raw()

	for _, r := range g.MustRunners(true) {
		r = g.MustRunner(r.Identifier)
		fmt.Println(r.Format(f))
	}
}
