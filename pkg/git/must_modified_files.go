package git

import "github.com/funtimecoding/soil/pkg/errors"

func MustModifiedFiles(path string) []string {
	result, e := ModifiedFiles(path)
	errors.PanicOnError(e)

	return result
}
