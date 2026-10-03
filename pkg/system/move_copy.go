package system

import "github.com/funtimecoding/soil/pkg/errors"

func MoveCopy(
	source string,
	destination string,
) {
	input := Open(source)
	defer errors.LogClose(input)
	output := Create(destination)
	defer errors.LogClose(output)
	Copy(input, output)

	if IsExecutable(source) {
		Executable(destination)
	}

	closeAndRemove(input, source)
}
