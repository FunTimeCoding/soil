package packager

import (
	"archive/tar"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
)

func addFileToTar(
	w *tar.Writer,
	filePath string,
	nameInArchive string,
) {
	f := system.Open(filePath)
	defer errors.PanicClose(f)
	s := system.FileStat(f)
	system.TarWriteHeader(
		w,
		&tar.Header{
			Name:    nameInArchive,
			Size:    s.Size(),
			Mode:    int64(s.Mode()),
			ModTime: s.ModTime(),
		},
	)
	system.Copy(f, w)
}
