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
	s.Add(Extract(p))

	for _, i := range p.Imports() {
		s.collect(i, seen)
	}
}
