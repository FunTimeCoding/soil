package face

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func Extract(p *types.Package) []*fact.Interface {
	return scopeInterfaces(p.Path(), p.Scope())
}
