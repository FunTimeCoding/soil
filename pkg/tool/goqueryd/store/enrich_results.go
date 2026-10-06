package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search"
)

func (s *Store) EnrichResults(
	results []search.Result,
	metadata map[string]string,
) []search.Result {
	var keys []record.DocumentKey

	for _, r := range results {
		keys = append(
			keys,
			record.DocumentKey{Collection: r.Collection, Path: r.Path},
		)
	}

	identifierMap := s.documentIdentifiers(keys)
	var identifiers []int

	for _, k := range keys {
		identifiers = append(identifiers, identifierMap[k])
	}

	s.enrichMetadata(results, identifiers)
	var enriched []search.Result

	for _, r := range results {
		if !matchesMetadata(r.Metadata, r.SourceType, metadata) {
			continue
		}

		enriched = append(enriched, r)
	}

	return enriched
}
