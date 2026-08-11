package scan

import (
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"go/ast"
)

func parseWebFile(
	v *virtual_file_system.System,
	path string,
) *ast.File {
	if !v.Has(path) {
		return nil
	}

	f, _, e := parse.Source(path, v.ReadString(path))

	if e != nil {
		return nil
	}

	return f
}
