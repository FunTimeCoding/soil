package symbol

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func decorated(
	set *token.FileSet,
	decorations *decoration.Set,
	owner *packages.Package,
	file *ast.File,
	node ast.Node,
) (dst.Node, error) {
	if _, e := decorations.DecorateFile(set, owner, file); e != nil {
		return nil, e
	}

	d := decorations.DecoratedNode(owner, node)

	if d == nil {
		return nil, validation.New(
			"no decorated node at %s",
			set.Position(node.Pos()),
		)
	}

	return d, nil
}
