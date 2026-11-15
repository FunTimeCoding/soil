package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_embedder"
	"testing"
)

func NewEmbedder(t *testing.T) *mock_embedder.Embedder {
	t.Helper()
	result := mock_embedder.New()
	t.Cleanup(result.Close)

	return result
}
