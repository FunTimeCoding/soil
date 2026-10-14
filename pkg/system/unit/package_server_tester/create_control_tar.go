package package_server_tester

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
)

func CreateControlTar(
	workDirectory string,
	controlDirectory string,
) {
	p := system.Open(
		filepath.Join(controlDirectory, constant.AlpineMetadataFile),
	)
	defer errors.PanicClose(p)
	segment := system.TarSegment(
		func(t *tar.Writer) {
			errors.PanicOnError(
				t.WriteHeader(
					&tar.Header{
						Name: constant.AlpineMetadataFile,
						Size: system.FileStat(p).Size(),
						Mode: 0644,
					},
				),
			)
			system.Copy(p, t)
		},
	)
	f := system.Create(filepath.Join(workDirectory, constant.AlpineControlFile))
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	system.Copy(bytes.NewReader(segment), z)
}
