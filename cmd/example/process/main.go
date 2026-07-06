package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/process"
)

func main() {
	for _, p := range process.New().Processes() {
		fmt.Printf(
			"%d %d %s\n",
			p.Pid(),
			p.PPid(),
			p.Executable(),
		)
	}
}
