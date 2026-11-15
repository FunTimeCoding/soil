package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"testing"
)

func NewReranker(t *testing.T) *mock_reranker.Reranker {
	t.Helper()
	result := mock_reranker.New()
	t.Cleanup(func() { errors.PanicClose(result) })

	return result
}
