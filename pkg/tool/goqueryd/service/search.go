package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Service) Search(o *search_option.Option) *result.Outcome {
	o.Reranker = s.reranker
	var outcome *result.Outcome

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
			return result.NewDegradedOutcome(e)
		}

		filtered := store.ExcludePaths(results, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		outcome = result.NewOutcome(s.store.EnrichResults(filtered, o.Metadata))
	} else {
		outcome = s.store.SearchWithFallback(o, s.embedder)
	}

	outcome.Facets = result.ComputeFacets(outcome.Results, 20)

	return outcome
}
