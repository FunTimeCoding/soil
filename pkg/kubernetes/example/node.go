package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
)

func Node() {
	k := client.NewEnvironment()
	f := constant.Format

	for _, n := range k.Nodes() {
		console.Line(n.Format(f))
	}
}
