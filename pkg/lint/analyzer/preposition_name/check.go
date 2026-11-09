package preposition_name

import (
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
	faces *face.Set,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		tests := testNames(
			file,
			strings.HasSuffix(
				p.Fset.Position(file.Pos()).Filename,
				soil.TestSuffix,
			),
		)
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				if i, okay := n.(*ast.Ident); okay && !tests[i] {
					checkName(p, results, faces, i)
				}

				return true
			},
		)
	}
}
