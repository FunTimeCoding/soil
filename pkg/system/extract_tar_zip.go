package system

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/join"
	"io"
)

func ExtractTarZip(
	sourceFile string,
	destinationDirectory string,
) {
	f := Open(sourceFile)
	defer errors.LogClose(f)
	z := GnuZipReader(f)
	defer errors.LogClose(z)
	r := tar.NewReader(z)

	for {
		h, e := r.Next()

		if e == io.EOF {
			break
		}

		errors.PanicOnError(e)
		path := join.Absolute(destinationDirectory, h.Name)

		switch h.Typeflag {
		case tar.TypeDir:
			MakeDirectory(path)
		case tar.TypeReg:
			output := OpenFile(path, HeaderFileMode(h))
			Copy(r, output)
			errors.PanicClose(output)
		}
	}
}
