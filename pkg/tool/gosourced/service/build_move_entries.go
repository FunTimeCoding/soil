package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
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
) ([]*relocation.Entry, string) {
	batch := make(map[string]bool)

	for _, name := range symbols {
		batch[name] = true
	}

	var result []*relocation.Entry

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
			newName = FlipName(symbol)
			flipped = true
		}

		node := ast.Node(declaration)

		if spec != nil {
			node = spec
		}

		name := targetFile

		if name == "" {
			name = fmt.Sprintf("%s.go", camel.ToSnake(newName))
		}

		entry := relocation.NewEntry()
		entry.Symbol = symbol
		entry.NewName = newName
		entry.Flipped = flipped
		entry.Object = o
		entry.File = file
		entry.Declaration = declaration
		entry.Spec = spec
		entry.Node = node
		entry.Carried = imports.UsedBy(file, node)
		entry.TargetFile = name
		result = append(result, entry)
	}

	return result, ""
}
