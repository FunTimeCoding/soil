package packager

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
)

func (p *Packager) CreateArchive() string {
	path := filepath.Join(p.WorkDirectory, constant.AlpineArchiveFile)
	p.writeArchive(path)

	return p.calculateArchiveHash(path)
}
