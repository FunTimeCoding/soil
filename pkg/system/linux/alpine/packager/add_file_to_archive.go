package packager

import (
	"archive/tar"
	"crypto/sha1"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func (p *Packager) addFileToArchive(
	w *tar.Writer,
	path string,
	i os.FileInfo,
) {
	h := system.TarHeader(i, "")
	h.Name = filepath.ToSlash(system.RelativePath(p.ArchiveDirectory, path))
	h.Format = tar.FormatPAX

	if !i.Mode().IsRegular() {
		system.TarWriteHeader(w, h)

		return
	}

	f := system.Open(path)
	a := sha1.New()
	system.Copy(f, a)
	errors.PanicClose(f)
	h.PAXRecords = map[string]string{
		"APK-TOOLS.checksum.SHA1": hex.EncodeToString(a.Sum(nil)),
	}
	system.TarWriteHeader(w, h)
	f = system.Open(path)
	defer errors.PanicClose(f)
	system.Copy(f, w)
}
