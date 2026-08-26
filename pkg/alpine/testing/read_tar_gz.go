package testing

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"io"
	"strings"
)

func ReadTarGz(payload []byte) map[string][]byte {
	gzr, e := gzip.NewReader(strings.NewReader(string(payload)))
	errors.PanicOnError(e)
	defer errors.PanicClose(gzr)
	t := tar.NewReader(gzr)
	result := make(map[string][]byte)

	for {
		header, f := t.Next()

		if f == io.EOF {
			break
		}

		errors.PanicOnError(f)

		if header.Typeflag == tar.TypeReg {
			content, g := io.ReadAll(t)
			errors.PanicOnError(g)
			result[header.Name] = content
		}
	}

	return result
}
