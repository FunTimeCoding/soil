package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/system"
)

func main() {
	w := system.WorkDirectory()
	fmt.Printf("Directory: %s\n", w)

	if !runtime.RunningFromBinary() {
		fmt.Println("Run from source")
	} else {
		fmt.Println("Run from binary")
	}
}
