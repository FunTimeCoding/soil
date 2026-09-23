package value_return

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isGeneratedType(
	p *packages.Package,
	named *types.Named,
) bool {
	filename := p.Fset.Position(named.Obj().Pos()).Filename

	for _, file := range p.Syntax {
		if p.Fset.File(file.Pos()).Name() == filename && ast.IsGenerated(file) {
			return true
		}
	}

	return lint.IsGeneratedFile(filename)
}
