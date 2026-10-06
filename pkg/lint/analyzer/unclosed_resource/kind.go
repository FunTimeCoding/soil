package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"golang.org/x/tools/go/packages"
)

func Kind() *index.Kind {
	return index.NewKind(
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
