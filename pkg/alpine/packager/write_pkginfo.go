package packager

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func (p *Packager) WritePKGINFO(archiveHash string) {
	content := renderPKGINFO(
		p.ExecutableName,
		p.PackageVersion,
		p.calculateInstalledSize(),
		archiveHash,
	)
	pkginfoPath := filepath.Join(p.ControlDirectory, constant.MetadataFile)
	errors.PanicOnError(os.WriteFile(pkginfoPath, []byte(content), 0644))
}
