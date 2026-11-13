package reference

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/load"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result/references"
)

func (s *Service) FindReferences(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
) (*output.Results, *references.References, error) {
	r := output.NewResultsWithDirectory(directory)
	pattern := packagePath

	if s.workspace.Full() {
		pattern = goModConstant.AllPackages
	}

	all, set, e := load.Packages(directory, pattern)

	if e != nil {
		return nil, nil, e
	}

	declaration, _, e := lookup.FindDeclaration(
		all,
		packagePath,
		symbol,
		receiver,
	)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return r, nil, nil
	}

	if s.workspace.Full() {
		locations := referenceLocations(directory, all, set, declaration, "")

		return r, references.New(symbol, locations), nil
	}

	locations := indexedLocations(
		directory,
		s.workspace.Workspace(directory).References(),
		all,
		set,
		packagePath,
		declaration,
		"",
	)

	return r, references.New(symbol, locations), nil
}
