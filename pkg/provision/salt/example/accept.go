package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Accept() {
	c := salt.NewEnvironment()
	result, e := c.AcceptKey("test-minion")

	if e != nil {
		console.Format("accept error: %v\n", e)

		return
	}

	console.Format("accepted: %v\n", result)
}
