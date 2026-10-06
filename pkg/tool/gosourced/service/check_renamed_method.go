package service

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
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

	w := s.workspace(directory)

	for _, i := range face.FromWorkspace(w, all).Satisfied(f) {
		addPositionConcern(
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
