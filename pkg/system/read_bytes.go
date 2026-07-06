package system

import "github.com/funtimecoding/soil/pkg/errors"

func ReadBytes(
	base string,
	name string,
) []byte {
	result, e := Root(base).ReadFile(name)
	errors.PanicOnError(e)

	return result
}
