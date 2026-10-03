package installed

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"path"
)

func (b *Binary) MainPackage(module string) string {
	if b.Package != "" && b.Package != constant.CommandLineArguments {
		return b.Package
	}

	return path.Join(module, constant.CommandDirectory, b.Name)
}
