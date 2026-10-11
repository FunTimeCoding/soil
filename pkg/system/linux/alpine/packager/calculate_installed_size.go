package packager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func (p *Packager) calculateInstalledSize() int64 {
	var result int64
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

				if i.Mode().IsRegular() {
					result += i.Size()
				}

				return nil
			},
		),
	)

	return result
}
