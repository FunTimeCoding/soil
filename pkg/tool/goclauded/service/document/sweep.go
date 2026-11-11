package document

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"maps"
	"slices"
)

func (s *Service) sweep(
	collection string,
	indexed map[string]client.SearchResult,
	current map[string]bool,
	named map[string]bool,
) (int, []error) {
	swept := 0
	var failed []error

	for _, path := range slices.Sorted(maps.Keys(indexed)) {
		r := indexed[path]

		if current[path] || !named[indexedSession(&r)] {
			continue
		}

		if e := s.deleteDocument(collection, path); e != nil {
			failed = append(failed, e)

			if connection.Classify(e) != nil {
				return swept, failed
			}

			continue
		}

		swept++
	}

	return swept, failed
}
