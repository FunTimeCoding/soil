package system

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/errors"
)

func TarWriteHeader(
	w *tar.Writer,
	h *tar.Header,
) {
	errors.PanicOnError(w.WriteHeader(h))
}
