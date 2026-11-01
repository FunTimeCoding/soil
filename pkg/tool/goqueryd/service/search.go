package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Service) Search(
	x context.Context,
	o *search_option.Option,
) (*result.Outcome, error) {
	o.Reranker = s.reranker
	var found []result.Search

	if o.Mode == "keyword" {
		results, e := s.store.SearchKeyword(
			o.Query,
			o.Limit+len(o.Exclude),
			o.Collection,
			o.Full,
			o.Metadata,
		)

		if e != nil {
			return nil, e
		}

		filtered := store.ExcludePaths(results, o.Exclude)
		found = s.store.EnrichResults(
			filtered[:min(o.Limit, len(filtered))],
			o.Metadata,
		)
	} else {
		vectors, e := s.embedTexts(x, []string{s.asQuery(o.Query)})

		if e != nil {
			return nil, e
		}

		found, e = s.store.SearchHybrid(o, vectors[0])

		if e != nil {
			return nil, e
		}
	}

	outcome := result.NewOutcome(found)
	outcome.Facets = result.ComputeFacets(found, 20)

	return outcome, nil
}
