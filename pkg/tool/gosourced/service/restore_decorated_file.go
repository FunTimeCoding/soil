package service

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func restoreDecoratedFile(
	resolver *resolve.Names,
	packagePath string,
	aliases map[string]string,
	file *dst.File,
	path string,
	out *sink.Sink,
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

	out.Write(path, buffer.Bytes())

	return nil
}
