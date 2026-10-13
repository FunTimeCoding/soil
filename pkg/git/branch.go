package git

import "github.com/funtimecoding/soil/pkg/errors"

func Branch(path string) string {
	b, e := MustOpen(path).Head()
	errors.PanicOnError(e)

	return b.Name().Short()
}
