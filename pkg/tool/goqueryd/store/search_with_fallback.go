package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Store) SearchWithFallback(
	o *search_option.Option,
	m face.Embedder,
) *result.Outcome {
	status, e := s.Status()

	if e != nil {
		return result.NewDegradedOutcome(e)
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
			return result.NewDegradedOutcome(f)
		}

		filtered := ExcludePaths(results, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		result := result.NewOutcome(filtered)
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
			return result.NewDegradedOutcome(g)
		}

		filtered := ExcludePaths(keyword, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		result := result.NewDegradedOutcome(f)
		result.Results = filtered

		return result
	}

	return result.NewOutcome(results)
}
