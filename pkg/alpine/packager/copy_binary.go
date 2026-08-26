package packager

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func (p *Packager) CopyBinary() {
	binDirectory := filepath.Join(p.ArchiveDirectory, "usr", "bin")
	system.MakeDirectory(binDirectory)
	destination := filepath.Join(binDirectory, p.ExecutableName)
	copyFile(p.ExecutablePath, destination)
	errors.PanicOnError(os.Chmod(destination, 0755))
}
