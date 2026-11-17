package service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_embedder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	m := mock_embedder.New()
	t.Cleanup(m.Close)
	a := mock_reranker.New()
	t.Cleanup(func() { errors.PanicClose(a) })
	c := mock_reranker.New()
	t.Cleanup(func() { errors.PanicClose(c) })
	result := &Tester{
		Embedder:   m,
		Reranker:   a,
		Counter:    c,
		Reporter:   memory.New(),
		Path:       filepath.Join(t.TempDir(), "goqueryd.db"),
		EmbedModel: constant.NomicEmbedText.Name,
		embedder:   m,
		reranker:   a,
		counter:    c,
		t:          t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
