package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Jobs() {
	c := salt.NewEnvironment()
	result, e := c.Jobs()

	if e != nil {
		fmt.Printf("error: %v\n", e)

		return
	}

	for _, j := range result {
		fmt.Printf("%s %s %s\n", j.JID, j.Function, j.StartTime)
	}
}
