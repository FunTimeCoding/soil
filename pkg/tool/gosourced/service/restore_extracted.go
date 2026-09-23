package service

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/decorator/resolver/gopackages"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
)

func restoreExtracted(
	directory string,
	file *dst.File,
	path string,
	dryRun bool,
) error {
	restorer := decorator.NewRestorerWithImports(
		constant.StandalonePath,
		gopackages.WithConfig(
			filepath.Dir(path),
			packages.Config{BuildFlags: resolve.BuildFlags(directory)},
		),
	)
	var buffer bytes.Buffer

	if e := restorer.Fprint(&buffer, file); e != nil {
		return e
	}

	if dryRun {
		return nil
	}

	return os.WriteFile(path, buffer.Bytes(), 0644)
}
