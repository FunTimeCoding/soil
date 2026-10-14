package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"golang.org/x/tools/go/packages"
)

func NewSummaries(
	loaded []*packages.Package,
	external ...*fact.Summary,
) *Summaries {
	summaries := make([]*fact.Summary, 0, len(loaded)+len(external))

	for _, p := range loaded {
		summaries = append(summaries, Extract(p))
	}

	return FromFacts(append(summaries, external...))
}
