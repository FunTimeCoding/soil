package store

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"

func (s *Store) enrichMetadata(
	results []SearchResult,
	identifiers []int,
) {
	metadata := s.metadataByDocuments(identifiers)

	for i := range results {
		results[i].Metadata = metadata[identifiers[i]]

		if results[i].Metadata != nil {
			results[i].SourceType = results[i].Metadata[constant.SourceType]
		}

		if results[i].SourceType == "" {
			results[i].SourceType = s.ResolveSourceType(
				results[i].Collection,
				results[i].Path,
			)
		}
	}
}
