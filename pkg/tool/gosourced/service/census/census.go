package census

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/result"
	"slices"
	"strings"
)

func (s *Service) Census(
	directory string,
	packagePath string,
	typeName string,
) (*result.Census, error) {
	loaded, _, e := s.workspace.CensusPackages(directory, packagePath)

	if e != nil {
		return nil, e
	}

	home := homePackage(loaded, packagePath)

	if home == nil {
		return nil, validation.New("unknown package: %s", packagePath)
	}

	subject := censusType(home, typeName)

	if subject == nil {
		return nil, validation.New(
			"no struct type %s in %s",
			typeName,
			packagePath,
		)
	}

	module := home.Module.Path
	methods, functions := declarations(home, subject, module)
	keys := map[string]bool{}

	for _, r := range methods {
		throughFunctions(r, functions)

		for k := range r.keys {
			keys[k] = true
		}
	}

	stores, f := storeReadings(directory, module, keys)

	if f != nil {
		return nil, f
	}

	tables, g := tableNames(directory, stores)

	if g != nil {
		return nil, g
	}

	countCallers(loaded, packagePath, subject, module, methods)
	names := make([]string, 0, len(methods))

	for n := range methods {
		names = append(names, n)
	}

	slices.SortFunc(
		names,
		func(a, b string) int {
			return strings.Compare(strings.ToLower(a), strings.ToLower(b))
		},
	)
	var entries []*result.CensusMethod

	for _, n := range names {
		entries = append(entries, methodResult(methods, n, stores, tables))
	}

	return result.NewCensus(
		packagePath,
		typeName,
		structFields(subject),
		entries,
	), nil
}
