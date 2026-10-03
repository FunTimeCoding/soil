package gofix

import (
	"bytes"
	"github.com/dave/dst/decorator"
	"go/parser"
	"go/token"
)

func formatPass(
	name string,
	source []byte,
	collapse bool,
) ([]*formatChange, []byte, error) {
	fileSet := token.NewFileSet()
	file, e := parser.ParseFile(fileSet, name, source, parser.ParseComments)

	if e != nil {
		return nil, nil, e
	}

	d := decorator.NewDecorator(fileSet)
	destination, e := d.DecorateFile(file)

	if e != nil {
		return nil, nil, e
	}

	changes := walkFormatEdits(destination, d, fileSet, source, collapse)

	if len(changes) == 0 {
		return nil, source, nil
	}

	var buffer bytes.Buffer

	if e = decorator.Fprint(&buffer, destination); e != nil {
		return nil, nil, e
	}

	return changes, buffer.Bytes(), nil
}
