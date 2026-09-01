package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Jobs() {
	c := salt.NewEnvironment()
	result, e := c.Jobs()

	if e != nil {
		console.Format("error: %v\n", e)

		return
	}

	for _, j := range result {
		console.Format("%s %s %s\n", j.JID, j.Function, j.StartTime)
	}
}
