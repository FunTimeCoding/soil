package index

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func typeDeclarations(p *packages.Package) map[string][]string {
	result := make(map[string][]string)

	for _, f := range p.Syntax {
		generated := ast.IsGenerated(f)

		for _, d := range f.Decls {
			g, okay := d.(*ast.GenDecl)

			if !okay || g.Tok != token.TYPE {
				continue
			}

			for _, s := range g.Specs {
				t := s.(*ast.TypeSpec)

				if !t.Name.IsExported() {
					continue
				}

				name := t.Name.Name

				if generated {
					result[name] = append(
						result[name],
						join.Space(name, "generated"),
					)
				}

				group := t.Doc

				if group == nil && !g.Lparen.IsValid() {
					group = g.Doc
				}

				if group == nil {
					continue
				}

				for _, c := range group.List {
					result[name] = append(
						result[name],
						join.Space(name, "comment", c.Text),
					)
				}
			}
		}
	}

	return result
}
