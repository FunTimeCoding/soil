package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/provision/salt"
)

func Accept() {
	c := salt.NewEnvironment()
	result, e := c.AcceptKey("test-minion")

	if e != nil {
		fmt.Printf("accept error: %v\n", e)

		return
	}

	fmt.Printf("accepted: %v\n", result)
}
