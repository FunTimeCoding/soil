package git

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"io"
)

func blobContent(
	r *git.Repository,
	h plumbing.Hash,
) ([]byte, error) {
	b, e := r.BlobObject(h)

	if e != nil {
		return nil, e
	}

	reader, f := b.Reader()

	if f != nil {
		return nil, f
	}

	defer errors.LogClose(reader)

	return io.ReadAll(reader)
}
