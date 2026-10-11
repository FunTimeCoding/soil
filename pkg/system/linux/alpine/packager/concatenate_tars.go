package packager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"os"
	"path/filepath"
)

func (p *Packager) ConcatenateTars() {
	controlPath := filepath.Join(p.WorkDirectory, constant.AlpineControlFile)
	archivePath := filepath.Join(p.WorkDirectory, constant.AlpineArchiveFile)
	outFile, e := os.Create(p.OutputFile)
	errors.PanicOnError(e)
	defer errors.PanicClose(outFile)
	appendFile(outFile, controlPath)
	appendFile(outFile, archivePath)
}
