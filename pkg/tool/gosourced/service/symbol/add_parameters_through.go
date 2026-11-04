package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/addition"
)

func (s *Service) addParametersThrough(
	directory string,
	additions []*addition.Parameter,
	imports []string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	var paths []string

	for _, a := range additions {
		paths = append(paths, a.PackagePath)
	}

	all, set, e := s.workspace.CensusPackages(directory, paths...)

	if e != nil {
		return nil, e
	}

	var sources []string

	for _, a := range additions {
		sources = append(sources, a.Type, a.Argument)
	}

	typing := extendGraph(directory, all, sources, imports)
	index := packageIndex(typing)
	explicit, f := importQualifiers(imports, index)

	if f != nil {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, f.Error(), "", false),
		)

		return r, nil
	}

	names := lookup.PackageNames(typing)
	targets := additionTargets(all, set, additions, names, explicit, index, r)

	if commit.HasUnfixed(r) {
		return r, nil
	}

	calls := additionCalls(all, set, targets, names, explicit, index, r)

	if commit.HasUnfixed(r) {
		return r, nil
	}

	decorations := decoration.NewSet()
	e = applyAdditions(set, decorations, targets, calls, r)

	if e != nil {
		return nil, e
	}

	e = decoration.Restore(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	return r, nil
}
