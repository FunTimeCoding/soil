package system

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/join"
	"os"
	"path/filepath"
)

func CreateTarZip(
	sourceDirectory string,
	destinationFile string,
) {
	f := Create(destinationFile)
	defer errors.LogClose(f)
	z := gzip.NewWriter(f)
	defer errors.LogClose(z)
	w := tar.NewWriter(z)
	defer errors.LogClose(w)
	errors.PanicOnError(
		filepath.Walk(
			sourceDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				errors.PanicOnError(e)
				h := TarHeader(i, i.Name())
				h.Name = filepath.ToSlash(
					join.Absolute(sourceDirectory, path),
				)
				TarWriteHeader(w, h)

				if !i.Mode().IsRegular() {
					return nil
				}

				l := Open(path)
				Copy(l, w)
				errors.PanicClose(l)

				return nil
			},
		),
	)
}
