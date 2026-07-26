package go_mod

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"golang.org/x/mod/modfile"
)

func Read() *modfile.File {
	result, e := modfile.Parse(
		constant.ModFile,
		system.ReadBytes(system.WorkDirectory(), constant.ModFile),
		nil,
	)
	errors.PanicOnError(e)

	return result
}
