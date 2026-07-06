package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
)

func RemotesRaw(path string) []*git.Remote {
	result, e := Open(path).Remotes()
	errors.FatalOnError(e)

	return result
}
