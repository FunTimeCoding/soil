package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/storer"
)

func TagsIterator(r *git.Repository) storer.ReferenceIter {
	result, e := r.Tags()
	errors.PanicOnError(e)

	return result
}
