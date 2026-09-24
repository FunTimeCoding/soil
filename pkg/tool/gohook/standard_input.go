package gohook

import (
	"github.com/funtimecoding/soil/pkg/system"
	"os"
)

func standardInput() string {
	if system.IsTerminal() {
		return ""
	}

	return string(system.ReadAll(os.Stdin))
}
