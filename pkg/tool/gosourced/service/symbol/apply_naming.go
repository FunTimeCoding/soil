package symbol

import (
	"github.com/dave/dst"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"go/token"
)

func applyNaming(
	set *token.FileSet,
	decorations *decoration.Set,
	sites []*naming.Site,
	name string,
	r *output.Results,
) error {
	for _, s := range sites {
		d, e := decorated(set, decorations, s.Owner, s.File, s.Signature)

		if e != nil {
			return e
		}

		fields := d.(*dst.FuncType).Params.List

		if s.Index < 0 {
			for i, field := range fields {
				blank := "_"

				if i == s.Field {
					blank = name
				}

				field.Names = []*dst.Ident{dst.NewIdent(blank)}
			}
		} else {
			fields[s.Field].Names[s.Index].Name = name
		}

		concern_line.Add(
			r,
			set,
			s.Signature.Pos(),
			constant.ConcernNamed,
			true,
			"%s: named %s",
			s.Label,
			name,
		)
	}

	return nil
}
