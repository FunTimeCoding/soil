package packages

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) movePackageFollow(
	directory string,
	packagePath string,
	targetPackagePath string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.workspace.ImportersLoad(directory, packagePath, true)

	if e != nil {
		return nil, e
	}

	decorations := decoration.NewSet()
	e = commit.RetargetImports(
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

	if name := lookup.ImportedName(all, packagePath); name != "" {
		names.Override(targetPackagePath, name)
	}

	return r, decoration.Restore(decorations, names, nil, out)
}
