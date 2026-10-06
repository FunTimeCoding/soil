package face

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Kind() *index.Kind {
	return index.NewKind(
		constant.InterfacesKind,
		func() any {
			return new([]*fact.Interface)
		},
		func(p *packages.Package) any {
			return new(Extract(p.Types))
		},
		func(p *types.Package) any {
			return new(Extract(p))
		},
	)
}
