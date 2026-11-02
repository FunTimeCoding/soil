package system

import "github.com/funtimecoding/soil/pkg/errors"

func ReadBytes(
	base string,
	name string,
) []byte {
	r := Root(base)
	defer errors.PanicClose(r)
	result, e := r.ReadFile(name)
	errors.PanicOnError(e)

	return result
}
