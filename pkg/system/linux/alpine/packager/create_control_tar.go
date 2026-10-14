package packager

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
)

func (p *Packager) CreateControlTar() {
	segment := system.TarSegment(
		func(t *tar.Writer) {
			addFileToTar(
				t,
				filepath.Join(p.ControlDirectory, constant.AlpineMetadataFile),
				constant.AlpineMetadataFile,
			)
		},
	)
	f := system.Create(
		filepath.Join(p.WorkDirectory, constant.AlpineControlFile),
	)
	defer errors.PanicClose(f)
	z := gzip.NewWriter(f)
	defer errors.PanicClose(z)
	system.Copy(bytes.NewReader(segment), z)
}
