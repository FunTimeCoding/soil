package service

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/guess"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"os"
)

func restoreExtracted(
	file *dst.File,
	path string,
) error {
	restorer := decorator.NewRestorerWithImports(
		constant.StandalonePath,
		guess.New(),
	)
	var buffer bytes.Buffer

	if e := restorer.Fprint(&buffer, file); e != nil {
		return e
	}

	return os.WriteFile(path, buffer.Bytes(), 0644)
}
