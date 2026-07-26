package decoration

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"os"
)

func WriteFile(
	file *dst.File,
	path string,
) error {
	var buffer bytes.Buffer

	if e := decorator.Fprint(&buffer, file); e != nil {
		return e
	}

	return os.WriteFile(path, buffer.Bytes(), 0644)
}
