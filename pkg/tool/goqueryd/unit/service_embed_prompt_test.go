package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/search_option"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"testing"
)

func TestDocumentIsEmbeddedWithItsModelsDocumentPrompt(t *testing.T) {
	m := service_tester.NewEmbedder(t)
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
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
	m := service_tester.NewEmbedder(t)
	v := service_tester.NewService(t, m, service_tester.NewReranker(t))
	_, e := v.Search(t.Context(), search_option.New("zulu", 10))
	assert.FatalOnError(t, e)
	assert.Strings(t, []string{"search_query: zulu"}, m.Inputs())
}

func TestUnknownEmbedModelNamesKnownOnes(t *testing.T) {
	o := service_tester.NewOption(
		t,
		service_tester.NewEmbedder(t),
		service_tester.NewReranker(t),
	)
	o.EmbedModel = "unknown"
	assert.PanicContains(
		t,
		`unknown embed model "unknown", known: nomic-embed-text, text-embedding-nomic-embed-text-v1.5@f16, qwen3-embedding:0.6b`,
		func() { service_tester.New(t, o, memory.New()) },
	)
}
