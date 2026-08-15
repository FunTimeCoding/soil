package gofix

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
)

func printDestinationDiffs(changed map[string]*dst.File) {
	for path, file := range changed {
		var buffer bytes.Buffer

		if e := decorator.Fprint(&buffer, file); e != nil {
			errors.Printf("print %s: %s\n", path, e)

			continue
		}

		original, e := os.ReadFile(path)

		if e != nil {
			errors.Printf("read %s: %s\n", path, e)

			continue
		}

		printDiff(path, original, buffer.Bytes())
	}
}
