package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"golang.org/x/tools/go/packages"
)

func NewSummaries(loaded []*packages.Package) *Summaries {
	summaries := make([]*fact.Summary, 0, len(loaded))

	for _, p := range loaded {
		summaries = append(summaries, Extract(p))
	}

	return FromFacts(summaries)
}
