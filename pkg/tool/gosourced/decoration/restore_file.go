package decoration

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func RestoreFile(
	resolver *resolve.Names,
	packagePath string,
	aliases map[string]string,
	file *dst.File,
	path string,
	out *sink.Sink,
) error {
	content, e := Restored(resolver, packagePath, aliases, file)

	if e != nil {
		return e
	}

	out.Write(path, content)

	return nil
}
