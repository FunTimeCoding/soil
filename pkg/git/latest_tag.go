package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5/plumbing"
	"time"
)

func LatestTag(path string) string {
	r := Open(path)
	var result string
	var latest time.Time
	errors.PanicOnError(
		TagsIterator(r).ForEach(
			func(f *plumbing.Reference) error {
				c := CommitFromHash(r, f.Hash())

				if c == nil {
					// Skip
					return nil
				}

				if c.Committer.When.After(latest) {
					latest = c.Committer.When
					result = f.Name().Short()
				}

				return nil
			},
		),
	)

	return result
}
