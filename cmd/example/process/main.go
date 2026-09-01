package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system/process"
)

func main() {
	for _, p := range process.New().Processes() {
		console.Format("%d %d %s\n", p.Pid(), p.PPid(), p.Executable())
	}
}
