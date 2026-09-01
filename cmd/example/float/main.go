package main

import "github.com/funtimecoding/soil/pkg/console"

func main() {
	f := 1.52
	console.Format("As f: %f\n", f)
	console.Format("As f.1: %.1f\n", f)
	console.Format("As f.2: %.2f\n", f)
	console.Format("As v: %v\n", f)
	console.Format("As +v: %+v\n", f)
}
