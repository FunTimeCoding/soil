package symbol

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/receiver_plan"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func renameReceiver(
	decorations *decoration.Set,
	all []*packages.Package,
	set *token.FileSet,
	p *receiver_plan.Plan,
	r *output.Results,
) error {
	owner, file := lookup.FindOwningFile(all, p.Method.Pos())

	if file == nil {
		return validation.New("no file holds %s", p.Label)
	}

	if _, e := decorations.DecorateFile(set, owner, file); e != nil {
		return e
	}

	for _, i := range p.Idents {
		d := decorations.DecoratedIdent(owner, i)

		if d == nil {
			return validation.New(
				"no decorated identifier at %s",
				set.Position(i.Pos()),
			)
		}

		d.Name = p.New
	}

	concern_line.Add(
		r,
		set,
		p.Method.Pos(),
		constant.ConcernRenamed,
		true,
		"%s: %s → %s, %d uses",
		p.Label,
		p.Old,
		p.New,
		len(p.Idents)-1,
	)

	return nil
}
