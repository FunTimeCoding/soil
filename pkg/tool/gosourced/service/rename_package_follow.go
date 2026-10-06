package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
	"path"
)

func (s *Service) renamePackageFollow(
	directory string,
	packagePath string,
	newName string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.importersLoad(directory, packagePath, true)

	if e != nil {
		return nil, e
	}

	oldName := importedName(all, packagePath)
	qualifiers, taken := collectPackageQualifiers(
		all,
		set,
		packagePath,
		oldName,
		newName,
	)

	if taken != "" {
		return failValidation(r, taken)
	}

	targetPackagePath := path.Join(path.Dir(packagePath), newName)
	decorations := decoration.NewSet()
	e = retargetImports(
		r,
		decorations,
		set,
		all,
		packagePath,
		targetPackagePath,
	)

	if e != nil {
		return nil, e
	}

	e = decorateQualifiers(r, decorations, set, qualifiers, oldName, newName)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(targetPackagePath, newName)

	return r, restoreDecorations(decorations, names, nil, out)
}
