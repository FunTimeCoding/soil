package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Store) SearchWithFallback(
	o *search_option.Option,
	m face.Embedder,
) *search.Outcome {
	status, e := s.Status()

	if e != nil {
		return search.NewDegradedOutcome(e)
	}

	if status.TotalEmbeddings == 0 {
		limit := o.Limit + len(o.Exclude)
		results, f := s.SearchKeyword(
			o.Query,
			limit,
			o.Collection,
			o.Full,
			o.Metadata,
		)

		if f != nil {
			return search.NewDegradedOutcome(f)
		}

		filtered := ExcludePaths(results, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		result := search.NewOutcome(filtered)
		result.Degraded = true

		return result
	}

	results, f := s.SearchHybrid(o, m)

	if f != nil {
		limit := o.Limit + len(o.Exclude)
		keyword, g := s.SearchKeyword(
			o.Query,
			limit,
			o.Collection,
			o.Full,
			o.Metadata,
		)

		if g != nil {
			return search.NewDegradedOutcome(g)
		}

		filtered := ExcludePaths(keyword, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		result := search.NewDegradedOutcome(f)
		result.Results = filtered

		return result
	}

	return search.NewOutcome(results)
}
