package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"io/fs"
)

func readPresent(
	base string,
	name string,
) ([]byte, bool) {
	r := system.Root(base)
	defer errors.PanicClose(r)
	result, e := r.ReadFile(name)

	if errors.Is(e, fs.ErrNotExist) {
		return nil, false
	}

	errors.PanicOnError(e)

	return result, true
}
