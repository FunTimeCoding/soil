package package_server_tester

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func CreateControlTar(
	workDirectory string,
	controlDirectory string,
) {
	f := system.Create(filepath.Join(workDirectory, constant.ControlFile))
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	p := system.Open(filepath.Join(controlDirectory, constant.MetadataFile))
	defer errors.PanicClose(p)
	errors.PanicOnError(
		t.WriteHeader(
			&tar.Header{
				Name: constant.MetadataFile,
				Size: system.FileStat(p).Size(),
				Mode: 0644,
			},
		),
	)
	system.Copy(p, t)
	errors.PanicFlush(t)
}
