package pattern

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func packageStructs(p *packages.Package) ([]string, []string) {
	scope := p.Types.Scope()
	var structs, receivers []string

	for _, name := range scope.Names() {
		t, okay := scope.Lookup(name).(*types.TypeName)

		if !okay || t.IsAlias() {
			continue
		}

		if file := lookup.SyntaxFileAt(p, t.Pos()); file != nil && ast.IsGenerated(file) {
			continue
		}

		named, structure := lookup.NamedStruct(t)

		if structure == nil {
			continue
		}

		structs = append(structs, name)

		if named.NumMethods() > 0 {
			receivers = append(receivers, name)
		}
	}

	return structs, receivers
}
