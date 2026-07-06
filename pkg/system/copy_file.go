package system

import "github.com/funtimecoding/soil/pkg/errors"

func CopyFile(
	source string,
	destination string,
) {
	s := Open(source)
	defer errors.LogClose(s)
	d := Create(destination)
	defer errors.LogClose(d)
	Copy(s, d)
}
