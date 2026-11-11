package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/symbol/relocator"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Service) moveMethodsFollow(
	directory string,
	packagePath string,
	typeName string,
	methods []string,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	objects := func(loaded []*packages.Package, _ *token.FileSet) []types.Object {
		return relocator.MethodObjects(loaded, packagePath, typeName, methods)
	}
	all, set, e := s.workspace.ReferenceLoad(
		directory,
		packagePath,
		"",
		objects,
	)

	if e != nil {
		return nil, e
	}

	for _, o := range objects(all, set) {
		for _, f := range resolve.FindAllReferences(all, o) {
			position := set.Position(f.Ident.Pos())

			if !system.InsideDirectory(directory, position.Filename) {
				continue
			}

			return commit.FailValidation(
				r,
				fmt.Sprintf(
					"%s is called from a replacing module at %s:%d",
					o.Name(),
					position.Filename,
					position.Line,
				),
			)
		}
	}

	return r, nil
}
