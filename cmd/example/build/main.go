package main

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system"
)

func main() {
	fmt.Printf("Executable: %s\n", system.ExecutablePath())
}
