package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
)

func (s *Service) Rechunk() ([]string, error) {
	documents, e := s.store.DocumentBodies("")

	if e != nil {
		return nil, e
	}

	stored, f := s.store.EmbeddingPositions()

	if f != nil {
		return nil, f
	}

	first := map[string]*record.DocumentBody{}

	for _, d := range documents {
		if x, okay := first[d.Hash]; !okay || d.Path < x.Path {
			first[d.Hash] = d
		}
	}

	var result []string

	for hash, d := range first {
		positions := stored[hash]

		if len(positions) == 0 ||
			samePositions(positions, chunk.Document(d.Body, d.Path, s.reranker)) {
			continue
		}

		if g := s.store.DeleteEmbeddings(hash); g != nil {
			return result, g
		}

		result = append(result, join.Slash([]string{d.Collection, d.Path}))
	}

	return result, nil
}
