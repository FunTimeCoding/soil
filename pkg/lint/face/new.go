package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"golang.org/x/tools/go/packages"
)

func New(loaded []*packages.Package) *Set {
	result := &Set{
		byMethod: make(map[string][]*fact.Interface),
		known:    make(map[string]bool),
	}
	seen := make(map[string]bool)

	for _, p := range loaded {
		if p.Types != nil {
			result.collect(p.Types, seen)
		}
	}

	return result
}
