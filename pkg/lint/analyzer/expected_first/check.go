package expected_first

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

// Check flags assert helpers whose expected parameter does not
// directly follow t - helpers mirror the assert shape, expected
// first, so the expected-slot analyzers read the right argument.
func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Recv != nil ||
				!strings.HasPrefix(f.Name.Name, constant.AssertHelperPrefix) {
				continue
			}

			if !takesTestingFirst(p, f) {
				continue
			}

			index := expectedIndex(f.Type.Params)

			if index <= 1 {
				continue
			}

			results.AddConcern(
				concern.NewFile(
					"expected_first",
					fmt.Sprintf(
						"assert helper %s: expected parameter must follow t",
						f.Name.Name,
					),
					p.Fset.Position(f.Pos()).Filename,
					false,
				),
			)
		}
	}
}
