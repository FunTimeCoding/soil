package system

import (
	"archive/zip"
	"github.com/funtimecoding/soil/pkg/errors"
	"path/filepath"
)

func ExtractZipFile(
	f *zip.File,
	directory string,
) {
	r, e := f.Open()
	errors.PanicOnError(e)
	defer errors.PanicClose(r)
	o := Create(filepath.Join(directory, f.Name))
	defer errors.PanicClose(o)
	Copy(r, o)
}
