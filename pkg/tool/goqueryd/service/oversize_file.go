package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize_chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service/oversize_file"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/record"
)

func (s *Service) oversizeFile(
	d *record.DocumentBody,
	allowance int,
) *oversize_file.File {
	var result *oversize_file.File

	for i, c := range chunk.Document(d.Body, d.Path, s.reranker) {
		tokens := s.reranker.Count(c.Text)

		if tokens <= allowance {
			continue
		}

		if result == nil {
			result = oversize_file.New(d.Collection, d.Path)
		}

		first, last := lineRange(d.Body, c)
		result.Add(oversize_chunk.New(i, first, last, len(c.Text), tokens))
	}

	return result
}
