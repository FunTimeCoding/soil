package packages

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
)

func (s *Service) renamePackageClauseFollow(
	directory string,
	packagePath string,
	newName string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.workspace.ImportersLoad(directory, packagePath, false)

	if e != nil {
		return nil, e
	}

	oldName := lookup.ImportedName(all, packagePath)
	qualifiers, taken := collectPackageQualifiers(
		all,
		set,
		packagePath,
		oldName,
		newName,
	)

	if taken != "" {
		return commit.FailValidation(r, taken)
	}

	decorations := decoration.NewSet()
	e = decorateQualifiers(r, decorations, set, qualifiers, oldName, newName)

	if e != nil {
		return nil, e
	}

	names := resolve.NewNames(all)
	names.Override(packagePath, newName)

	return r, decoration.Restore(decorations, names, nil, out)
}
