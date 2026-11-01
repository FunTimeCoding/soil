package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
)

func (s *Service) overWindow(chunks []chunk.Chunk) []error {
	result := make([]error, len(chunks))
	allowance := s.counter.Allowance()

	for i, c := range chunks {
		if tokens := s.counter.Count(s.asDocument(c.Text)); tokens > allowance {
			result[i] = validation.New(
				constant.EmbedOverWindow,
				tokens,
				allowance,
			)
		}
	}

	return result
}
