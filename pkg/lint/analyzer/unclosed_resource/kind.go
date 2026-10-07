package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index/kind"
	"golang.org/x/tools/go/packages"
)

func Kind() *kind.Kind {
	return kind.New(
		constant.SummaryKind,
		func() any {
			return fact.NewSummary()
		},
		func(p *packages.Package) any {
			return Extract(p)
		},
		nil,
	)
}
