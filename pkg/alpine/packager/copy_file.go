package packager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"os"
)

func copyFile(
	source string,
	destination string,
) {
	sourceFile, e := os.Open(source)
	errors.PanicOnError(e)
	defer errors.PanicClose(sourceFile)
	destinationFile, f := os.Create(destination)
	errors.PanicOnError(f)
	defer errors.PanicClose(destinationFile)
	_, g := io.Copy(destinationFile, sourceFile)
	errors.PanicOnError(g)
	errors.PanicOnError(destinationFile.Sync())
}
