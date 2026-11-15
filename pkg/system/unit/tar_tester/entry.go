package tar_tester

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/system"
)

func Entry(
	name string,
	content string,
) func(*tar.Writer) {
	return func(w *tar.Writer) {
		system.TarWriteHeader(
			w,
			&tar.Header{Name: name, Size: int64(len(content)), Mode: 0644},
		)
		system.TarWrite(w, []byte(content))
	}
}
