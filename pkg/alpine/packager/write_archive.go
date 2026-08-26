package packager

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func (p *Packager) writeArchive(path string) {
	f := system.Create(path)
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	defer errors.PanicClose(t)
	errors.PanicOnError(
		filepath.Walk(
			p.ArchiveDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				if e != nil {
					return e
				}

				if path == p.ArchiveDirectory {
					return nil
				}

				p.addFileToArchive(t, path, i)

				return nil
			},
		),
	)
}
