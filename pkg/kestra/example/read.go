package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kestra"
)

func Read() {
	k := kestra.NewEnvironment()

	if true {
		fmt.Println("Namespaces")

		for _, n := range k.Namespaces() {
			fmt.Printf("Namespace: %+v\n", n)

			for _, f := range k.Flows(n) {
				fmt.Printf("  Flow: %+v\n", f)
			}
		}
	}

	if false {
		// 404
		fmt.Println("Users")

		for _, f := range k.Users() {
			fmt.Printf("  %+v\n", f)
		}
	}
}
