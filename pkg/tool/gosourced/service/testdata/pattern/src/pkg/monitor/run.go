package monitor

import (
	"example/pkg/client"
	"fmt"
)

func Run(c *client.Client) {
	if c.Ready() {
		fmt.Println("ready")
	}

	if c.Steady() {
		fmt.Println("steady")
	}

	fmt.Println("done", 2)
}
