package unit

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/system"
)

func writeTarEntry(name string) func(*tar.Writer) {
	return func(w *tar.Writer) {
		system.TarWriteHeader(
			w,
			&tar.Header{Name: name, Size: int64(len(name)), Mode: 0644},
		)
		system.TarWrite(w, []byte(name))
	}
}
