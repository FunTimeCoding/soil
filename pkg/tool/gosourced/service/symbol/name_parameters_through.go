package symbol

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/naming"
	"go/token"
)

func (s *Service) nameParametersThrough(
	directory string,
	request *naming.Request,
	paths []string,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	refuse := func(message string) (*output.Results, error) {
		r.AddConcern(
			concern.NewFile(constant.ConcernValidation, message, "", false),
		)

		return r, nil
	}

	if !token.IsIdentifier(request.Name) || request.Name == "_" {
		return refuse(fmt.Sprintf("%s is not a parameter name", request.Name))
	}

	if len(paths) == 0 {
		return refuse(fmt.Sprintf("no package matches %s", request.PackagePath))
	}

	all, set, e := s.workspace.CensusPackages(directory, paths...)

	if e != nil {
		return nil, e
	}

	typing := extendGraph(
		directory,
		all,
		[]string{request.Type},
		request.Imports,
	)
	index := packageIndex(typing)
	explicit, f := importQualifiers(request.Imports, index)

	if f != nil {
		return refuse(f.Error())
	}

	scope := map[string]bool{}

	for _, p := range paths {
		scope[p] = true
	}

	sites := namingSites(
		all,
		set,
		scope,
		request,
		lookup.PackageNames(typing),
		explicit,
		index,
		r,
	)

	if commit.HasUnfixed(r) {
		return r, nil
	}

	decorations := decoration.NewSet()
	e = applyNaming(set, decorations, sites, request.Name, r)

	if e != nil {
		return nil, e
	}

	e = decoration.Restore(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	return r, nil
}
