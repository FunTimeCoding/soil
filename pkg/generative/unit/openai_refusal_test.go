package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestOpenAIRefusedBodyIsValidationError(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusBadRequest),
		upstream_server.WithAnswer(
			`{"error":{"message":"Invalid body: failed to parse JSON value","type":"invalid_request","code":"invalid_json"}}`,
		),
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "Invalid body", e.Error())
}

func TestOpenAIServerErrorIsUnexpectedAndNotRetried(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusInternalServerError),
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.True(t, unexpected.Is(e))
	assert.Integer(t, 1, s.Count())
}

func TestBusyOpenAIServerIsUnavailable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusServiceUnavailable),
		upstream_server.WithAnswer(`{"error":{"message":"busy"}}`),
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.True(t, unavailable.Is(e))
}
