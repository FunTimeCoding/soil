package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5/plumbing"
)

func Tags(path string) []string {
	var result []string
	tags, e := Open(path).Tags()
	errors.PanicOnError(e)
	errors.PanicOnError(
		tags.ForEach(
			func(reference *plumbing.Reference) error {
				name := reference.Name()

				if name.IsTag() {
					result = append(result, name.Short())
				}

				return nil
			},
		),
	)

	return result
}
