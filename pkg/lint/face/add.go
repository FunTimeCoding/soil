package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Set) Add(interfaces []*fact.Interface) {
	for _, i := range interfaces {
		k := join.Dot([]string{i.Package, i.Name})

		if s.known[k] {
			continue
		}

		s.known[k] = true

		for method := range i.Methods {
			s.byMethod[method] = append(s.byMethod[method], i)
		}
	}
}
