package service

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/sink"
)

func (s *Service) movePackageFollow(
	directory string,
	packagePath string,
	targetPackagePath string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.importersLoad(directory, packagePath, true)

	if e != nil {
		return nil, e
	}

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

	names := resolve.NewNames(all)

	if name := importedName(all, packagePath); name != "" {
		names.Override(targetPackagePath, name)
	}

	return r, restoreDecorations(decorations, names, nil, out)
}
