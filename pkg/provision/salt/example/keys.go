package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Keys() {
	c := salt.NewEnvironment()
	result, e := c.Keys()

	if e != nil {
		console.Format("error: %v\n", e)

		return
	}

	console.Format("accepted: %v\n", result.Minions)
	console.Format("pending: %v\n", result.MinionsPre)
	console.Format("denied: %v\n", result.MinionsDenied)
	console.Format("rejected: %v\n", result.MinionsRejected)
}
