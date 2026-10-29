package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/embed"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings/join"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnOpenAIClientSendsItsModel(t *testing.T) {
	var model string
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				model = sentModel(r)
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write(
					[]byte(
						`{"object":"list","data":[{"object":"embedding","index":0,"embedding":[0.5]}]}`,
					),
				)
				errors.PanicOnError(e)
			},
		),
	)
	defer s.Close()
	c := openai.NewBase(join.Empty(s.URL, "/v1"), "nomic-embed-text-v1.5.f16")
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.FatalOnError(t, e)
	assert.String(t, "nomic-embed-text-v1.5.f16", model)
}

func TestAnOllamaClientSendsItsModel(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"embeddings":[[0.5]]}`)
	host, port := upstream_tester.HostPort(t, s.URL)
	c := ollama.New(
		ollama.WithHost(host),
		ollama.WithPort(port),
		ollama.WithEmbedModel("nomic-embed-text:v1.5"),
	)
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.FatalOnError(t, e)
	assert.String(t, "nomic-embed-text:v1.5", sentModel(last()))
}

func TestAnOllamaClientWithoutAModelSendsTheDefault(t *testing.T) {
	s, last := upstream_tester.New(t, http.StatusOK, `{"embeddings":[[0.5]]}`)
	host, port := upstream_tester.HostPort(t, s.URL)
	c := ollama.New(ollama.WithHost(host), ollama.WithPort(port))
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.FatalOnError(t, e)
	assert.String(t, "nomic-embed-text", sentModel(last()))
}

func TestTheEmbedModelDefaultsPerScheme(t *testing.T) {
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

func TestTheEmbedModelEnvironmentOverridesTheDefault(t *testing.T) {
	t.Setenv(constant.EmbedTargetEnvironment, "openai://alfa:1234")
	t.Setenv(constant.EmbedModelEnvironment, "nomic-embed-text-v1.5.f16")
	assert.String(t, "nomic-embed-text-v1.5.f16", embed.ModelEnvironment())
}

func TestAnUnknownEmbedSchemeHasNoDefaultModel(t *testing.T) {
	assert.PanicContains(
		t,
		"unknown embed target scheme: bravo",
		func() { embed.DefaultModel("bravo") },
	)
}
