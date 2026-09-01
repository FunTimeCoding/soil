package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/kestra"
)

func Read() {
	k := kestra.NewEnvironment()

	if true {
		console.Line("Namespaces")

		for _, n := range k.Namespaces() {
			console.Format("Namespace: %+v\n", n)

			for _, f := range k.Flows(n) {
				console.Format("  Flow: %+v\n", f)
			}
		}
	}

	if false {
		// 404
		console.Line("Users")

		for _, f := range k.Users() {
			console.Format("  %+v\n", f)
		}
	}
}
