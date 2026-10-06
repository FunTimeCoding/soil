package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Service) Search(o *search_option.Option) *search.Outcome {
	o.Reranker = s.reranker
	var outcome *search.Outcome

	if o.Mode == "keyword" {
		limit := o.Limit + len(o.Exclude)
		results, e := s.store.SearchKeyword(
			o.Query,
			limit,
			o.Collection,
			o.Full,
			o.Metadata,
		)

		if e != nil {
			return search.NewDegradedOutcome(e)
		}

		filtered := store.ExcludePaths(results, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		outcome = search.NewOutcome(s.store.EnrichResults(filtered, o.Metadata))
	} else {
		outcome = s.store.SearchWithFallback(o, s.embedder)
	}

	outcome.Facets = search.ComputeFacets(outcome.Results, 20)

	return outcome
}
