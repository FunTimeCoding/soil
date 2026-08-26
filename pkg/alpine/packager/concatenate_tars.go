package packager

import (
	"github.com/funtimecoding/soil/pkg/alpine/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"path/filepath"
)

func (p *Packager) ConcatenateTars() {
	controlPath := filepath.Join(p.WorkDirectory, constant.ControlFile)
	archivePath := filepath.Join(p.WorkDirectory, constant.ArchiveFile)
	outFile, e := os.Create(p.OutputFile)
	errors.PanicOnError(e)
	defer errors.PanicClose(outFile)
	appendFile(outFile, controlPath)
	appendFile(outFile, archivePath)
}
