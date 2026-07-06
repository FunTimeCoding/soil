package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func Head(r *git.Repository) *plumbing.Reference {
	result, e := r.Head()
	errors.PanicOnError(e)

	return result
}
