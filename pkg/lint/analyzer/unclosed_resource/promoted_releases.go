package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func promotedReleases(
	p *packages.Package,
	summary *fact.Summary,
) {
	scope := p.Types.Scope()

	for _, name := range scope.Names() {
		t, okay := scope.Lookup(name).(*types.TypeName)

		if !okay || summary.Releases[typeKey(t)] != nil {
			continue
		}

		s, okay := t.Type().Underlying().(*types.Struct)

		if !okay {
			continue
		}

		_, index, _ := types.LookupFieldOrMethod(
			types.NewPointer(t.Type()),
			true,
			t.Pkg(),
			constant.CloseMethod,
		)

		if len(index) < 2 || !s.Field(index[0]).Embedded() {
			continue
		}

		summary.Releases[typeKey(t)] = []fact.Release{
			{Field: s.Field(index[0]).Name()},
		}
	}
}
