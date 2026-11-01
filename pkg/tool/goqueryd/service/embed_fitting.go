package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/chunk"
)

func (s *Service) embedFitting(
	x context.Context,
	chunks []chunk.Chunk,
	refusals []error,
) ([][]float32, error) {
	result := make([][]float32, len(chunks))
	var texts []string
	var fitting []int

	for i, c := range chunks {
		if refusals[i] == nil {
			texts = append(texts, s.asDocument(c.Text))
			fitting = append(fitting, i)
		}
	}

	if len(texts) == 0 {
		return result, nil
	}

	embeddings, e := s.embedTexts(x, texts)

	if e == nil {
		for k, i := range fitting {
			result[i] = embeddings[k]
		}

		return result, nil
	}

	if !validation.Is(e) {
		return nil, e
	}

	for _, i := range fitting {
		v, f := s.embedTexts(x, []string{s.asDocument(chunks[i].Text)})

		if f == nil {
			result[i] = v[0]

			continue
		}

		if !validation.Is(f) {
			return nil, f
		}

		refusals[i] = f
	}

	return result, nil
}
