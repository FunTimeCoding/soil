package unclosed_resource

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkTransfer(
	p *packages.Package,
	results *output.Results,
	s *Summaries,
	t resource_transfer.Transfer,
) {
	o := t.Holder.Origin().Obj()
	key := typeKey(o)

	if !inModule(p, o.Pkg()) && !s.knowsHolder(key) {
		return
	}

	held := t.Field.Type()

	if t.Container {
		held = elementType(held)
	}

	var message string

	switch index := closeIndex(t.Holder); {
	case held == nil || !carriesObligation(held):
		message = fmt.Sprintf(
			"%s.%s holds %s behind %s, which has no Close",
			o.Name(),
			t.Field.Name(),
			types.TypeString(
				p.TypesInfo.TypeOf(t.Value),
				types.RelativeTo(p.Types),
			),
			types.TypeString(t.Field.Type(), types.RelativeTo(p.Types)),
		)
	case index == nil:
		message = fmt.Sprintf(
			"%s holds %s but has no Close to release it",
			o.Name(),
			t.Field.Name(),
		)
	case len(index) > 1:
		if fieldName(t.Holder, index[0]) == t.Field.Name() {
			return
		}

		message = fmt.Sprintf(
			"%s.Close is promoted from %s and never releases %s",
			o.Name(),
			fieldName(t.Holder, index[0]),
			t.Field.Name(),
		)
	case s.releasesField(key, t.Field.Name()):
		return
	default:
		message = fmt.Sprintf(
			"%s.Close never releases %s",
			o.Name(),
			t.Field.Name(),
		)
	}

	results.AddConcern(
		concern.NewPosition(
			"unclosed_resource",
			message,
			p.Fset.Position(t.Value.Pos()),
		),
	)
}
