package packager

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"path/filepath"
)

func (p *Packager) CreateArchive() string {
	path := filepath.Join(p.WorkDirectory, constant.ArchiveFile)
	p.writeArchive(path)

	return p.calculateArchiveHash(path)
}
