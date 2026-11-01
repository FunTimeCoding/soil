package git

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func RevisionFiles(
	path string,
	revision string,
	keep func(string) bool,
) (map[string][]byte, error) {
	r, e := Open(path)

	if e != nil {
		return nil, e
	}

	scope, f := worktreeScope(r, path)

	if f != nil {
		return nil, f
	}

	h, g := r.ResolveRevision(plumbing.Revision(revision))

	if g != nil {
		return nil, not_found.New(constant.RevisionKind, revision)
	}

	c := CommitFromHash(r, *h)

	if c == nil {
		return nil, not_found.New(constant.RevisionKind, revision)
	}

	t, i := c.Tree()

	if i != nil {
		return nil, i
	}

	result := map[string][]byte{}
	j := t.Files().ForEach(
		func(o *object.File) error {
			name, inside := scoped(o.Name, scope)

			if !inside || !keep(name) {
				return nil
			}

			content, k := blobContent(r, o.Hash)

			if k != nil {
				return k
			}

			result[name] = content

			return nil
		},
	)

	if j != nil {
		return nil, j
	}

	return result, nil
}
