package face

import "go/types"

func (s *Set) collect(
	p *types.Package,
	seen map[string]bool,
) {
	if seen[p.Path()] {
		return
	}

	seen[p.Path()] = true
	scope := p.Scope()

	for _, name := range scope.Names() {
		typeName, isTypeName := scope.Lookup(name).(*types.TypeName)

		if !isTypeName {
			continue
		}

		f, isFace := typeName.Type().Underlying().(*types.Interface)

		if !isFace {
			continue
		}

		for i := range f.NumMethods() {
			method := f.Method(i).Name()
			s.byMethod[method] = append(s.byMethod[method], f)
		}
	}

	for _, i := range p.Imports() {
		s.collect(i, seen)
	}
}
