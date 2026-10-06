package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/types"
)

func Extract(p *types.Package) []*fact.Interface {
	var result []*fact.Interface
	scope := p.Scope()

	for _, name := range scope.Names() {
		typeName, isTypeName := scope.Lookup(name).(*types.TypeName)

		if !isTypeName {
			continue
		}

		f, isFace := typeName.Type().Underlying().(*types.Interface)

		if !isFace || f.NumMethods() == 0 {
			continue
		}

		methods := make(map[string]string, f.NumMethods())

		for i := range f.NumMethods() {
			methods[f.Method(i).Name()] = index.Signature(f.Method(i))
		}

		result = append(result, fact.NewInterface(p.Path(), name, methods))
	}

	return result
}
