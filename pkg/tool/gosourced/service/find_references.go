package service

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
)

func (s *Service) FindReferences(
	directory string,
	packagePath string,
	symbol string,
	receiver string,
) (*output.Results, *result.References, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := resolve.LoadPackages(directory, "./...")

	if e != nil {
		return nil, nil, e
	}

	declaration, _, e := findDeclaration(all, packagePath, symbol, receiver)

	if e != nil {
		r.AddConcern(
			concern.NewFile("validation", e.Error(), "", false),
		)

		return r, nil, nil
	}

	locations := referenceLocations(directory, all, set, declaration, "")

	return r, result.NewReferences(symbol, locations), nil
}
