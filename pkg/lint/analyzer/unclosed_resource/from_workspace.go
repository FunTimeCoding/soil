package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"github.com/funtimecoding/soil/pkg/source/index"
	"golang.org/x/tools/go/packages"
)

func FromWorkspace(
	w *index.Workspace,
	loaded []*packages.Package,
) *Summaries {
	live := make(map[string]bool, len(loaded))
	var summaries []*fact.Summary

	for _, p := range loaded {
		live[p.PkgPath] = true
		summaries = append(summaries, Extract(p))
	}

	for path, s := range index.Facts[*fact.Summary](w, Kind()) {
		if !live[path] {
			summaries = append(summaries, s)
		}
	}

	return FromFacts(summaries)
}
