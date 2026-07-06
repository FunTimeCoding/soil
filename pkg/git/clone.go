package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
)

func Clone(
	locator string,
	path string,
) *git.Repository {
	result, e := git.PlainClone(
		path,
		false,
		&git.CloneOptions{URL: locator},
	)
	errors.FatalOnError(e)

	return result
}
