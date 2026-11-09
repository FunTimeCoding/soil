package unguarded_peer

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func addClients(
	p *packages.Package,
	l *ast.FieldList,
	result map[types.Object]bool,
) {
	if l == nil {
		return
	}

	guard := false

	for _, f := range l.List {
		if isGuard(p.TypesInfo.TypeOf(f.Type)) {
			guard = true
		}
	}

	if !guard {
		return
	}

	for _, f := range l.List {
		if !isPeerClient(p.TypesInfo.TypeOf(f.Type)) {
			continue
		}

		for _, name := range f.Names {
			result[p.TypesInfo.Defs[name]] = true
		}
	}
}
