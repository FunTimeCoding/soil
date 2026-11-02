package packages

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func decorateModified(
	decorations *decoration.Set,
	set *token.FileSet,
	all []*packages.Package,
	modified map[string]*ast.File,
) error {
	for _, file := range modified {
		owner, syntax := lookup.FindOwningFile(all, file.Pos())

		if syntax == nil {
			continue
		}

		if _, e := decorations.DecorateFile(set, owner, syntax); e != nil {
			return e
		}
	}

	return nil
}
