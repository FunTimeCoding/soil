package decoration

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/resolve"
)

func Restored(
	resolver *resolve.Names,
	packagePath string,
	aliases map[string]string,
	file *dst.File,
) ([]byte, error) {
	r := decorator.NewRestorerWithImports(packagePath, resolver)
	f := r.FileRestorer()

	for importPath, alias := range aliases {
		f.Alias[importPath] = alias
	}

	var buffer bytes.Buffer

	if e := f.Fprint(&buffer, file); e != nil {
		return nil, e
	}

	return imports.JoinBlock(buffer.Bytes()), nil
}
