package string_constant

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func parseConstants(
	result map[string][]knownConstant,
	path string,
	p string,
) {
	set := token.NewFileSet()
	file, e := parser.ParseFile(set, path, nil, 0)

	if e != nil {
		return
	}

	if p == "" {
		p = file.Name.Name
	}

	for _, d := range file.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok != token.CONST {
			continue
		}

		for _, s := range g.Specs {
			v, okay := s.(*ast.ValueSpec)

			if !okay {
				continue
			}

			for i, name := range v.Names {
				if i >= len(v.Values) {
					continue
				}

				l, okay := v.Values[i].(*ast.BasicLit)

				if !okay || l.Kind != token.STRING {
					continue
				}

				value := strings.Trim(l.Value, "\"")
				result[value] = append(
					result[value],
					knownConstant{name: name.Name, packageName: p},
				)
			}
		}
	}
}
