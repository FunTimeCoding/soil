package symbol

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/rename_plan"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol_rename"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func (s *Service) planFamily(
	directory string,
	all []*packages.Package,
	set *token.FileSet,
	m *types.Func,
	n *symbol_rename.Rename,
	r *output.Results,
) []*rename_plan.Plan {
	members, refusals := methodFamily(all, m)
	slices.SortFunc(
		members,
		func(a *types.Func, b *types.Func) int {
			x := set.Position(a.Pos())
			y := set.Position(b.Pos())

			return cmp.Or(
				cmp.Compare(x.Filename, y.Filename),
				cmp.Compare(x.Line, y.Line),
			)
		},
	)

	for _, f := range members {
		position := set.Position(f.Pos())
		r.AddConcern(
			concern.NewLine(
				constant.ConcernFamily,
				types.ObjectString(f, nil),
				position.Filename,
				position.Line,
				"",
				true,
			),
		)
	}

	refusals = slices.Concat(
		refusals,
		familyOutside(
			face.FromWorkspace(s.workspace.Workspace(directory), all),
			all,
			members,
		),
		familyCollisions(all, members, n.NewName),
		familyGenerated(set, members),
	)

	for _, refusal := range refusals {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, refusal, "", false),
		)
	}

	if len(refusals) > 0 {
		return nil
	}

	result := make([]*rename_plan.Plan, 0, len(members))

	for _, f := range members {
		result = append(
			result,
			rename_plan.New(
				symbol_rename.New(
					n.OldName,
					n.NewName,
					receiverNamed(f).Obj().Name(),
				),
				f,
				resolve.FindAllReferences(all, f),
			))
	}

	return result
}
