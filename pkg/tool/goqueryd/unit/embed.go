package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"strings"
)

func (f *FixedEmbedder) Embed(
	_ context.Context,
	v []string,
) ([][]float32, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()
	f.calls++

	if f.embedding != nil {
		return nil, f.embedding
	}

	if f.failure != nil {
		return nil, f.failure
	}

	for _, s := range v {
		if f.refused != "" && strings.Contains(s, f.refused) {
			return nil, validation.New(
				"ollama: the input length exceeds the context length",
			)
		}
	}

	result := make([][]float32, len(v))

	for i := range v {
		result[i] = []float32{1, 0}
	}

	return result, nil
}
