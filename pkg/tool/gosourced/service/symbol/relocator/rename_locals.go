package relocator

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocation"
	"go/ast"
	"go/token"
	"sort"
)

func renameLocals(
	r *output.Results,
	decorations *decoration.Set,
	set *token.FileSet,
	q *relocation.FileQualification,
) {
	var declared []*ast.Ident

	for ident, letter := range q.Locals {
		if d := decorations.DecoratedIdent(q.Owner, ident); d != nil {
			d.Name = letter
		}

		if q.Owner.TypesInfo.Defs[ident] != nil {
			declared = append(declared, ident)
		}
	}

	sort.Slice(
		declared,
		func(
			i int,
			j int,
		) bool {
			return declared[i].Pos() < declared[j].Pos()
		},
	)

	for _, ident := range declared {
		concern_line.Add(
			r,
			set,
			ident.Pos(),
			constant.ConcernRenamed,
			true,
			"local %s → %s, the import keeps its name",
			ident.Name,
			q.Locals[ident],
		)
	}
}
