package match

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strings"
)

func Parse(source string) (*Pattern, error) {
	set := token.NewFileSet()
	file, e := parser.ParseFile(
		set,
		"pattern.go",
		join.Empty("package pattern\n\n", source),
		parser.SkipObjectResolution,
	)

	if e != nil {
		return nil, fmt.Errorf("pattern does not parse: %w", e)
	}

	for _, d := range file.Decls {
		declaration, okay := d.(*ast.FuncDecl)

		if !okay {
			continue
		}

		if declaration.Body == nil || len(declaration.Body.List) != 1 {
			return nil, fmt.Errorf(
				"pattern body must hold exactly one statement",
			)
		}

		holes := map[string]string{}

		if declaration.Type.Params != nil {
			for _, field := range declaration.Type.Params.List {
				var b strings.Builder
				errors.PanicOnError(printer.Fprint(&b, set, field.Type))

				for _, name := range field.Names {
					holes[name.Name] = b.String()
				}
			}
		}

		statement := declaration.Body.List[0]

		if f := checkSpreadHoles(statement, holes); f != nil {
			return nil, f
		}

		return &Pattern{Holes: holes, Statement: statement}, nil
	}

	return nil, fmt.Errorf("pattern must declare a function")
}
