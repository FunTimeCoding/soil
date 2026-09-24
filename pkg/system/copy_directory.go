package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func CopyDirectory(
	source string,
	destination string,
) {
	errors.PanicOnError(os.CopyFS(destination, os.DirFS(source)))
}
