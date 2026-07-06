package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
)

func Status(path string) git.Status {
	status, e := Tree(path).Status()
	errors.PanicOnError(e)

	return status
}
