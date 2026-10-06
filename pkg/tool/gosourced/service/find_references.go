package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result/references"
)

func (s *Service) FindReferences(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
) (*output.Results, *references.References, error) {
	r := output.NewResultsWithDirectory(directory)
	pattern := packagePath

	if s.full {
		pattern = "./..."
	}

	all, set, e := loadPackages(directory, pattern)

	if e != nil {
		return nil, nil, e
	}

	declaration, _, e := findDeclaration(all, packagePath, symbol, receiver)

	if e != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, e.Error(), "", false),
		)

		return r, nil, nil
	}

	if s.full {
		locations := referenceLocations(directory, all, set, declaration, "")

		return r, references.New(symbol, locations), nil
	}

	locations := indexedLocations(
		directory,
		s.workspace(directory).References(),
		all,
		set,
		packagePath,
		declaration,
		"",
	)

	return r, references.New(symbol, locations), nil
}
