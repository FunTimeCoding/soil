package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func ReplaceFile(
	source string,
	destination string,
) {
	s := Open(source)
	defer errors.LogClose(s)
	d, e := os.CreateTemp(filepath.Dir(destination), filepath.Base(destination))
	errors.PanicOnError(e)
	Copy(s, d)
	errors.PanicClose(d)
	ChangeMode(d.Name(), Mode(source))
	errors.PanicOnError(os.Rename(d.Name(), destination))
}
