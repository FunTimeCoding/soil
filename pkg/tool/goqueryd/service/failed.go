package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/failed_chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
)

func (s *Service) Failed(collection string) ([]*failed_chunk.Chunk, error) {
	failures, e := s.store.Failures(collection)

	if e != nil {
		return nil, e
	}

	result := make([]*failed_chunk.Chunk, len(failures))

	for i, f := range failures {
		result[i] = failed_chunk.New(f.Collection, f.Path, f.Sequence, f.Reason)
		chunks := chunk.Document(f.Body, f.Path, s.reranker)

		if f.Sequence < len(chunks) {
			c := chunks[f.Sequence]
			result[i].FirstLine, result[i].LastLine = lineRange(f.Body, c)
			result[i].Bytes = len(c.Text)
		}
	}

	return result, nil
}
