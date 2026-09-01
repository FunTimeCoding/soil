package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/system"
)

func main() {
	w := system.WorkDirectory()
	console.Format("Directory: %s\n", w)

	if !runtime.RunningFromBinary() {
		console.Line("Run from source")
	} else {
		console.Line("Run from binary")
	}
}
