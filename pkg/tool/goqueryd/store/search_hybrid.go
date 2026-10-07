package store

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
)

func (s *Store) SearchHybrid(
	o *search_option.Option,
	m face.Embedder,
) ([]result.Search, error) {
	fetchFull := o.Full || o.Reranker != nil
	keywordResults, e := s.SearchKeyword(
		o.Query,
		o.Limit*2,
		o.Collection,
		fetchFull,
		o.Metadata,
	)

	if e != nil {
		return nil, e
	}

	vectorResults, f := s.SearchVector(
		o.Query,
		o.Limit*2,
		o.Collection,
		fetchFull,
		o.Metadata,
		m,
	)

	if f != nil {
		return nil, f
	}

	merged, bodies := fuse(o.Exclude, keywordResults, vectorResults)
	candidates := rerankCandidates(
		o,
		merged[:min(o.Limit*3, len(merged))],
		bodies,
	)
	unenriched := make([]result.Search, 0, len(candidates))

	for _, c := range candidates {
		if !o.Full {
			c.Search.Body = ""
		}

		unenriched = append(unenriched, c.Search)
	}

	enriched := s.EnrichResults(unenriched, o.Metadata)

	return enriched[:min(o.Limit, len(enriched))], nil
}
