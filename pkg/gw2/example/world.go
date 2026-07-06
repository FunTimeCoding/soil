package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gw2"
)

func World() {
	fmt.Printf("Worlds: %+v\n", gw2.NewEnvironment().Worlds())
}
