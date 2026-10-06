package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path"
)

func (s *Service) moveSymbolsFollow(
	directory string,
	packagePath string,
	symbols []string,
	filePath string,
	targetPackagePath string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	objects := func(
		loaded []*packages.Package,
		loadedSet *token.FileSet,
	) []types.Object {
		return movedObjects(
			directory,
			loaded,
			loadedSet,
			packagePath,
			symbols,
			filePath,
		)
	}
	all, set, e := s.referenceLoad(directory, packagePath, "", objects)

	if e != nil {
		return nil, e
	}

	var external []relocation.QualifiedReference

	for _, o := range objects(all, set) {
		for _, f := range resolve.FindAllReferences(all, o) {
			position := set.Position(f.Ident.Pos())

			if !system.InsideDirectory(directory, position.Filename) {
				continue
			}

			external = append(
				external,
				relocation.QualifiedReference{Reference: f, NewName: o.Name()},
			)
		}
	}

	if len(external) == 0 {
		return r, nil
	}

	targetPackageName := importedName(all, targetPackagePath)

	if targetPackageName == "" {
		targetPackageName = path.Base(targetPackagePath)
	}

	qualifications, blocked := planQualifications(
		set,
		external,
		packagePath,
		targetPackagePath,
		targetPackageName,
	)

	if qualifications == nil {
		return failValidation(
			r,
			fmt.Sprintf("no available import name in %s", blocked),
		)
	}

	decorations := decoration.NewSet()
	e = applyQualifications(
		r,
		decorations,
		set,
		qualifications,
		targetPackagePath,
	)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(targetPackagePath, targetPackageName)

	return r, restoreDecorations(decorations, names, nil, out)
}
