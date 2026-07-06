package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func TagObject(
	r *git.Repository,
	h plumbing.Hash,
) *object.Tag {
	result, e := r.TagObject(h)
	errors.PanicOnError(e)

	return result
}
