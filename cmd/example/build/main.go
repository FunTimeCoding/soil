package main

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/system"
)

func main() {
	console.Format("Executable: %s\n", system.ExecutablePath())
}
