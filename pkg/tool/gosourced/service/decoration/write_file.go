package decoration

import (
	"bytes"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func WriteFile(
	file *dst.File,
	path string,
	out *sink.Sink,
) error {
	var buffer bytes.Buffer

	if e := decorator.Fprint(&buffer, file); e != nil {
		return e
	}

	out.Write(path, buffer.Bytes())

	return nil
}
