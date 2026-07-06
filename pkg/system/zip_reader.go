package system

import (
	"archive/zip"
	"github.com/funtimecoding/soil/pkg/errors"
)

func ZipReader(path string) *zip.ReadCloser {
	result, e := zip.OpenReader(path)
	errors.PanicOnError(e)

	return result
}
