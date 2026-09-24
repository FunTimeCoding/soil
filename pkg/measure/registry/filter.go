package registry

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/measure/language"
)

func (r *Registry) Filter(names []string) *Registry {
	if len(names) == 0 {
		return r
	}

	var selected []*language.Language

	for _, n := range names {
		l := r.ByName(n)

		if l == nil {
			panic(fmt.Sprintf("unknown language: %s", n))
		}

		selected = append(selected, l)
	}

	return New(selected...)
}
