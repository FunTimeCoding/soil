package packager

import (
	"archive/tar"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
)

func (p *Packager) CreateControlTar() {
	f := system.Create(
		filepath.Join(p.WorkDirectory, constant.AlpineControlFile),
	)
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	t := tar.NewWriter(z)
	addFileToTar(
		t,
		filepath.Join(p.ControlDirectory, constant.AlpineMetadataFile),
		constant.AlpineMetadataFile,
	)
	errors.PanicFlush(t)
}
