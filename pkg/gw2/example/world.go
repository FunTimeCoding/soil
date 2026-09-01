package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
)

func World() {
	console.Format("Worlds: %+v\n", gw2.NewEnvironment().Worlds())
}
