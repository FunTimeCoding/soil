package gofix

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func writeDestinationFiles(changed map[string]*dst.File) {
	for path, file := range changed {
		var buffer bytes.Buffer

		if e := decorator.Fprint(&buffer, file); e != nil {
			errors.Printf("print %s: %s\n", path, e)

			continue
		}

		if e := os.WriteFile(path, buffer.Bytes(), 0644); e != nil {
			errors.Printf("write %s: %s\n", path, e)
		}
	}
}
