package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/argument"
)

func main() {
	a := argument.NewSimple("argument-example")
	a.ParseSimple()
	fmt.Printf(
		"Positional argument 0: %s\n",
		a.Argument(0),
	)
}
