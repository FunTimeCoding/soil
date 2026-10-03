package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func CommitFromHash(
	r *git.Repository,
	h plumbing.Hash,
) *object.Commit {
	if t, tagFail := r.TagObject(h); tagFail == nil {
		if c, commitFail := t.Commit(); commitFail == nil {
			return c
		}
	}

	if c, objectFail := r.CommitObject(h); objectFail == nil {
		return c
	}

	return nil
}
