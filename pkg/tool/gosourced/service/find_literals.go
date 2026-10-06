package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/pattern_site"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/result"
)

func (s *Service) FindLiterals(
	directory string,
	packagePath string,
	name string,
) (*output.Results, *result.Literals, error) {
	r := output.NewResultsWithDirectory(directory)
	all, set, e := s.censusPackages(directory, packagePath)

	if e != nil {
		return nil, nil, e
	}

	declaration, _, f := findDeclaration(all, packagePath, name, "")

	if f != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, f.Error(), "", false),
		)

		return r, nil, nil
	}

	named, structure := namedStruct(declaration)

	if structure == nil {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernValidation,
				fmt.Sprintf("%s is not a struct type in %s", name, packagePath),
				"",
				false,
			),
		)

		return r, nil, nil
	}

	census := literalCensus(all, set, named, structure, packagePath)
	contents := map[string][]byte{}
	var entries []*pattern_site.Entry

	for _, site := range census.Sites {
		entry, g := literalEntry(directory, set, contents, site, site.Shape)

		if g != nil {
			return nil, nil, g
		}

		entries = append(entries, entry)
	}

	return r, result.NewLiterals(
		name,
		census.Total,
		census.Inside,
		census.Expected,
		groupEntries(entries),
	), nil
}
