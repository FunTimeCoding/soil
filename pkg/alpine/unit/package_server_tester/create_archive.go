package package_server_tester

import (
	"archive/tar"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func CreateArchive(
	workDirectory string,
	archiveDirectory string,
) string {
	path := filepath.Join(workDirectory, constant.ArchiveFile)
	f := system.Create(path)
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	defer errors.PanicClose(t)
	errors.PanicOnError(
		filepath.Walk(
			archiveDirectory,
			func(
				path string,
				i os.FileInfo,
				e error,
			) error {
				errors.PanicOnError(e)

				if path == archiveDirectory {
					return nil
				}

				relPath := system.RelativePath(archiveDirectory, path)
				h := system.TarHeader(i, "")
				h.Name = filepath.ToSlash(relPath)
				h.Format = tar.FormatPAX
				system.TarWriteHeader(t, h)

				if i.Mode().IsRegular() {
					file := system.Open(path)
					defer errors.PanicClose(file)
					system.Copy(file, t)
				}

				return nil
			},
		),
	)
	o := system.Open(path)
	defer errors.PanicClose(o)
	h := sha256.New()
	system.Copy(o, h)

	return hex.EncodeToString(h.Sum(nil))
}
