package system

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func MoveCopy(
	source string,
	destination string,
) {
	input := Open(source)
	defer errors.LogClose(input)
	output := Create(destination)
	defer errors.LogClose(output)
	Copy(input, output)
	// Close before trying to remove for Windows
	// https://stackoverflow.com/a/64943554/246801
	errors.PanicClose(input)

	if IsExecutable(source) {
		Executable(destination)
	}

	errors.PanicOnError(os.Remove(source))
}
