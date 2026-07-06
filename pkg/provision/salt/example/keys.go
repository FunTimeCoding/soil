package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Keys() {
	c := salt.NewEnvironment()
	result, e := c.Keys()

	if e != nil {
		fmt.Printf("error: %v\n", e)

		return
	}

	fmt.Printf("accepted: %v\n", result.Minions)
	fmt.Printf("pending: %v\n", result.MinionsPre)
	fmt.Printf("denied: %v\n", result.MinionsDenied)
	fmt.Printf("rejected: %v\n", result.MinionsRejected)
}
