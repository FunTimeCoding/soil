package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestDocumentIsEmbeddedWithItsModelsDocumentPrompt(t *testing.T) {
	q := service_tester.New(t)
	m := q.Embedder
	v := q.Service()
	v.MustPushDocument(
		t.Context(),
		"notes",
		"alfa.md",
		"# Alfa\n\nzulu.\n",
		nil,
	)
	assert.Strings(
		t,
		[]string{"search_document: # Alfa\n\nzulu.\n"},
		m.Inputs(),
	)
}

func TestQueryIsEmbeddedWithItsModelsQueryPrompt(t *testing.T) {
	q := service_tester.New(t)
	m := q.Embedder
	v := q.Service()
	_, e := v.Search(t.Context(), search_option.New("zulu", 10))
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"search_query: zulu"}, m.Inputs())
}

func TestUnknownEmbedModelNamesKnownOnes(t *testing.T) {
	q := service_tester.New(t, service_tester.WithEmbedModel("unknown"))
	assert.PanicContains(
		t,
		`unknown embed model "unknown", known: nomic-embed-text, text-embedding-nomic-embed-text-v1.5@f16, qwen3-embedding:0.6b`,
		func() { q.Service() },
	)
}
