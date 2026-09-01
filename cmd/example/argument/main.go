package main

import (
	"github.com/funtimecoding/soil/pkg/argument"
	"github.com/funtimecoding/soil/pkg/console"
)

func main() {
	a := argument.NewSimple("argument-example")
	a.ParseSimple()
	console.Format("Positional argument 0: %s\n", a.Argument(0))
}
