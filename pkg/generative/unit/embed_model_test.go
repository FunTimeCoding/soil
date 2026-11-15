package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/generative/unit/embedding_tester"
	"github.com/funtimecoding/soil/pkg/generative/unit/ollama_tester"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestOpenAIClientSendsItsModel(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"object":"list","data":[{"object":"embedding","index":0,"embedding":[0.5]}]}`,
		),
		upstream_server.WithObject(),
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), "nomic-embed-text-v1.5.f16")
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.FatalOnError(t, e)
	assert.String(t, "nomic-embed-text-v1.5.f16", embedding_tester.Model(s))
}

func TestOllamaClientSendsItsModel(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"embeddings":[[0.5]]}`),
	)
	_, e := ollama_tester.New(
		t,
		s.URL,
		ollama.WithEmbedModel("nomic-embed-text:v1.5"),
	).Embed(t.Context(), []string{"alfa"})
	assert.FatalOnError(t, e)
	assert.String(t, "nomic-embed-text:v1.5", embedding_tester.Model(s))
}

func TestOllamaClientWithoutModelSendsDefault(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"embeddings":[[0.5]]}`),
	)
	_, e := ollama_tester.New(t, s.URL).Embed(t.Context(), []string{"alfa"})
	assert.FatalOnError(t, e)
	assert.String(t, "nomic-embed-text", embedding_tester.Model(s))
}

func TestEmbedModelDefaultsPerScheme(t *testing.T) {
	t.Setenv(constant.EmbedModelEnvironment, "")
	t.Setenv(constant.EmbedTargetEnvironment, "ollama://alfa:11434")
	assert.String(t, "nomic-embed-text", embed.ModelEnvironment())
	t.Setenv(constant.EmbedTargetEnvironment, "openai://alfa:1234")
	assert.String(
		t,
		"text-embedding-nomic-embed-text-v1.5@f16",
		embed.ModelEnvironment(),
	)
}

func TestEmbedModelEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(constant.EmbedTargetEnvironment, "openai://alfa:1234")
	t.Setenv(constant.EmbedModelEnvironment, "nomic-embed-text-v1.5.f16")
	assert.String(t, "nomic-embed-text-v1.5.f16", embed.ModelEnvironment())
}

func TestUnknownEmbedSchemeHasNoDefaultModel(t *testing.T) {
	assert.PanicContains(
		t,
		"unknown embed target scheme: bravo",
		func() { embed.DefaultModel("bravo") },
	)
}
