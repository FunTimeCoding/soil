package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
)

func MustOpen(path string) *git.Repository {
	r, e := Open(path)
	errors.PanicOnError(e)

	return r
}
