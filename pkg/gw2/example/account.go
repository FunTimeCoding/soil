package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/gw2"
)

func Account() {
	console.Format("Account: %+v\n", gw2.NewEnvironment().Account())
}
