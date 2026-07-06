package go_mod

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/mod/modfile"
)

func Read() *modfile.File {
	result, e := modfile.Parse(
		ModFile,
		system.ReadBytes(system.WorkDirectory(), ModFile),
		nil,
	)
	errors.PanicOnError(e)

	return result
}
