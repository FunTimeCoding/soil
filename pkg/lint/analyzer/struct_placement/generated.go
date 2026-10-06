package struct_placement

import (
	"github.com/funtimecoding/soil/pkg/lint"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func generated(
	p *packages.Package,
	position token.Pos,
) bool {
	filename := p.Fset.Position(position).Filename

	if lint.IsGeneratedFile(filename) {
		return true
	}

	for _, file := range p.Syntax {
		if p.Fset.File(file.Pos()).Name() == filename {
			return ast.IsGenerated(file)
		}
	}

	return false
}
