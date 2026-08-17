package string_constant

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkConstant(
	p *packages.Package,
	results *output.Results,
	i *ast.Ident,
) {
	c, okay := p.TypesInfo.Uses[i].(*types.Const)

	if !okay || c.Pkg() == nil {
		return
	}

	if !strings.ContainsRune(strings.Split(c.Pkg().Path(), "/")[0], '.') {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"string_constant",
			fmt.Sprintf(
				"constant %s.%s in expected value should be a literal",
				c.Pkg().Name(),
				c.Name(),
			),
			p.Fset.Position(i.Pos()).Filename,
			false,
		),
	)
}
