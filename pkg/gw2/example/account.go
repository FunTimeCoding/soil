package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gw2"
)

func Account() {
	fmt.Printf("Account: %+v\n", gw2.NewEnvironment().Account())
}
