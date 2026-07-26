package service

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"os"
)

func restoreDecoratedFile(
	resolver *resolve.Names,
	packagePath string,
	aliases map[string]string,
	file *dst.File,
	path string,
) error {
	r := decorator.NewRestorerWithImports(packagePath, resolver)
	f := r.FileRestorer()

	for importPath, alias := range aliases {
		f.Alias[importPath] = alias
	}

	var buffer bytes.Buffer

	if e := f.Fprint(&buffer, file); e != nil {
		return e
	}

	return os.WriteFile(path, buffer.Bytes(), 0644)
}
