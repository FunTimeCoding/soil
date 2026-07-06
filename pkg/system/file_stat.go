package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func FileStat(f *os.File) os.FileInfo {
	result, e := f.Stat()
	errors.PanicOnError(e)

	return result
}
