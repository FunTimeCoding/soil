package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CommitObject(
	r *git.Repository,
	h plumbing.Hash,
) *object.Commit {
	result, e := r.CommitObject(h)
	errors.PanicOnError(e)

	return result
}
