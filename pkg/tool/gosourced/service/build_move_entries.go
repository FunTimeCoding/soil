package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"unicode"
)

func buildMoveEntries(
	all []*packages.Package,
	p *packages.Package,
	packagePath string,
	symbols []string,
	targetFile string,
) ([]*moveEntry, string) {
	batch := make(map[string]bool)

	for _, name := range symbols {
		batch[name] = true
	}

	var result []*moveEntry

	for _, symbol := range symbols {
		o, _, e := findDeclaration(all, packagePath, symbol, "")

		if e != nil {
			return nil, e.Error()
		}

		file, declaration, spec := findDeclarationNode(p, o)

		if declaration == nil {
			return nil, fmt.Sprintf("declaration not found: %s", symbol)
		}

		if v, okay := spec.(*ast.ValueSpec); okay && len(v.Names) > 1 {
			for _, n := range v.Names {
				if !batch[n.Name] {
					return nil, fmt.Sprintf(
						"%s shares a spec with %s - move them together",
						symbol,
						n.Name,
					)
				}
			}

			if targetFile == "" {
				return nil, fmt.Sprintf(
					"%s is declared in a multi-name spec - pass target_file",
					symbol,
				)
			}
		}

		newName := symbol
		flipped := false

		if unicode.IsLower(rune(symbol[0])) {
			newName = flipName(symbol)
			flipped = true
		}

		node := ast.Node(declaration)

		if spec != nil {
			node = spec
		}

		name := targetFile

		if name == "" {
			name = fmt.Sprintf("%s.go", toSnakeCase(newName))
		}

		result = append(
			result,
			&moveEntry{
				symbol:      symbol,
				newName:     newName,
				flipped:     flipped,
				object:      o,
				file:        file,
				declaration: declaration,
				spec:        spec,
				node:        node,
				carried:     imports.UsedBy(file, node),
				targetFile:  name,
			},
		)
	}

	return result, ""
}
