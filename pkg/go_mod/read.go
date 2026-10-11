package go_mod

import (
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/mod/modfile"
)

func Read() *modfile.File {
	return ReadDirectory(system.WorkDirectory())
}
