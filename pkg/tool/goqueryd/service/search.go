package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Service) Search(o *search_option.Option) *store.SearchOutcome {
	o.Reranker = s.reranker
	var outcome *store.SearchOutcome

	if o.Mode == "keyword" {
		limit := o.Limit + len(o.Exclude)
		results, e := s.store.SearchKeyword(
			o.Query,
			limit,
			o.Collection,
			o.Full,
		)

		if e != nil {
			return store.NewDegradedSearchOutcome(e)
		}

		filtered := store.ExcludePaths(results, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		outcome = store.NewSearchOutcome(
			s.store.EnrichResults(filtered, o.Metadata),
		)
	} else {
		outcome = s.store.SearchWithFallback(o, s.ollama)
	}

	outcome.Facets = store.ComputeFacets(outcome.Results, 20)

	return outcome
}
