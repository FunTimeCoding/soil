package packager

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"path/filepath"
)

func (p *Packager) CreateControlTar() {
	f := system.Create(filepath.Join(p.WorkDirectory, constant.ControlFile))
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	addFileToTar(
		t,
		filepath.Join(p.ControlDirectory, constant.MetadataFile),
		constant.MetadataFile,
	)
	errors.PanicFlush(t)
}
