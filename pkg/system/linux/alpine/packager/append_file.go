package packager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
)

func appendFile(
	destination *os.File,
	sourcePath string,
) {
	f := system.Open(sourcePath)
	defer errors.PanicClose(f)
	system.Copy(f, destination)
}
