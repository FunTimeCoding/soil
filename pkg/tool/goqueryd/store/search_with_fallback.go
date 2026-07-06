package store

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Store) SearchWithFallback(
	o *search_option.Option,
	l *ollama.Client,
) *SearchOutcome {
	status, e := s.Status()

	if e != nil {
		return &SearchOutcome{Degraded: true, Cause: e}
	}

	if status.TotalEmbeddings == 0 {
		limit := o.Limit + len(o.Exclude)
		results, f := s.SearchKeyword(
			o.Query,
			limit,
			o.Collection,
			o.Full,
		)

		if f != nil {
			return &SearchOutcome{Degraded: true, Cause: f}
		}

		filtered := ExcludePaths(results, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		return &SearchOutcome{
			Results:  filtered,
			Degraded: true,
		}
	}

	results, f := s.SearchHybrid(o, l)

	if f != nil {
		limit := o.Limit + len(o.Exclude)
		keyword, g := s.SearchKeyword(
			o.Query,
			limit,
			o.Collection,
			o.Full,
		)

		if g != nil {
			return &SearchOutcome{Degraded: true, Cause: g}
		}

		filtered := ExcludePaths(keyword, o.Exclude)

		if len(filtered) > o.Limit {
			filtered = filtered[:o.Limit]
		}

		return &SearchOutcome{
			Results:  filtered,
			Degraded: true,
			Cause:    f,
		}
	}

	return &SearchOutcome{Results: results}
}
