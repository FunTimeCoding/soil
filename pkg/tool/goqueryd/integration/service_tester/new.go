package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	s := base.New(t)
	s.SkipUnreachable(t)
	t.Cleanup(s.Close)

	return &Tester{
		t:       t,
		Service: service.New(s.Store(), s.Embedder(), s.Reranker()),
	}
}
