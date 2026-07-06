package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/kubernetes/constant"
)

func Pod() {
	k := client.NewEnvironment()
	f := constant.Format

	for _, n := range k.Pods(nil) {
		fmt.Println(n.Format(f))
	}
}
