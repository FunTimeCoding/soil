package service_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	queryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_reranker"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/option"
	"path/filepath"
	"testing"
)

func NewOption(
	t *testing.T,
	m face.ReachableEmbedder,
	a queryd.Reranker,
) *option.Query {
	t.Helper()
	result := option.New()
	result.LitePath = filepath.Join(t.TempDir(), "goqueryd.db")
	result.EmbedTarget = constant.FixtureEmbedTarget
	result.EmbedModel = constant.FixtureEmbedModel
	result.Embedder = func() face.ReachableEmbedder { return m }
	result.Reranker = func() queryd.Reranker { return a }
	result.EmbedCounter = func() queryd.EmbedCounter { return mock_reranker.New() }

	return result
}
