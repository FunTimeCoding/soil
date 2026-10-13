package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"testing"
)

func TestModelByNameFindsKnownModel(t *testing.T) {
	m := rerank.ModelByName("gte-reranker-modernbert-base")
	assert.String(t, "gte-reranker-modernbert-base", m.Name)
	assert.Integer(t, 512, m.SequenceLength)
}

func TestModelByNameUnknownPanicsListingKnownModels(t *testing.T) {
	assert.PanicContains(
		t,
		`unknown rerank model "bge-reranker-large", known: bge-reranker-base, gte-reranker-modernbert-base`,
		func() { rerank.ModelByName("bge-reranker-large") },
	)
}
