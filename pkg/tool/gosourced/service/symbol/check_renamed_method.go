package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/concern_line"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Service) checkRenamedMethod(
	directory string,
	all []*packages.Package,
	set *token.FileSet,
	declaration types.Object,
	r *output.Results,
) {
	f, okay := declaration.(*types.Func)

	if !okay {
		return
	}

	receiver := f.Type().(*types.Signature).Recv()

	if receiver == nil {
		return
	}

	w := s.workspace.Workspace(directory)
	modules := workspaceModules(all)

	for _, i := range face.FromWorkspace(w, all).Satisfied(f) {
		if _, declares := i.Methods[f.Name()]; declares &&
			inWorkspace(modules, i.Package) {
			concern_line.Add(
				r,
				set,
				f.Pos(),
				constant.ConcernValidation,
				false,
				"%s satisfies %s.%s - renaming it breaks that interface; rename %s.%s to rename them together",
				f.Name(),
				i.Package,
				i.Name,
				i.Name,
				f.Name(),
			)

			continue
		}

		concern_line.Add(
			r,
			set,
			f.Pos(),
			constant.ConcernValidation,
			false,
			"%s satisfies %s.%s - renaming it breaks that interface",
			f.Name(),
			i.Package,
			i.Name,
		)
	}
}
