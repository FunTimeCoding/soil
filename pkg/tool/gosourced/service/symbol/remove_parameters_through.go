package symbol

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/commit"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/sink"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/removal"
	"go/ast"
)

func (s *Service) removeParametersThrough(
	directory string,
	removals []*removal.Parameter,
	out *sink.Sink,
) (*output.Results, error) {
	r := output.NewResultsWithDirectory(directory)
	var paths []string

	for _, p := range removals {
		paths = append(paths, p.PackagePath)
	}

	all, set, e := s.workspace.CensusPackages(directory, paths...)

	if e != nil {
		return nil, e
	}

	targets := removalTargets(all, removals, r)

	if commit.HasUnfixed(r) {
		return r, nil
	}

	calls := removalCalls(all, set, targets, r)
	var removed []ast.Node

	for _, c := range calls {
		for _, a := range c.Arguments {
			removed = append(removed, a)
		}
	}

	checkBodyUses(set, targets, removed, r)
	checkSideEffects(set, calls, r)

	if commit.HasUnfixed(r) {
		return r, nil
	}

	locals, removed := planLocals(all, set, calls, removed, r)

	if commit.HasUnfixed(r) {
		return r, nil
	}

	decorations := decoration.NewSet()
	e = applyRemovals(set, decorations, targets, calls, locals, r)

	if e != nil {
		return nil, e
	}

	reportLeftovers(all, set, targets, calls, locals, removed, r)
	e = decoration.Restore(decorations, resolve.NewNames(all), nil, out)

	if e != nil {
		return nil, e
	}

	return r, nil
}
