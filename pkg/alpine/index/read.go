package index

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"io"
	"os"
)

func Read(path string) ([]*Entry, error) {
	f, e := os.Open(path)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(f)
	z, g := gzip.NewReader(f)

	if g != nil {
		return nil, g
	}

	defer errors.LogClose(z)
	t := tar.NewReader(z)

	for {
		h, k := t.Next()

		if k == io.EOF {
			break
		}

		if k != nil {
			return nil, k
		}

		if h.Name != constant.IndexFile {
			continue
		}

		b, m := io.ReadAll(t)

		if m != nil {
			return nil, m
		}

		return parse(string(b)), nil
	}

	return nil, not_found.Format("no %s in %s", constant.IndexFile, path)
}
