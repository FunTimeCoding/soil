package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"sync/atomic"
	"testing"
)

func TestAnOpenAIRefusedBodyIsAValidationError(t *testing.T) {
	e := openaiEmbedding(
		t,
		http.StatusBadRequest,
		`{"error":{"message":"Invalid body: failed to parse JSON value","type":"invalid_request","code":"invalid_json"}}`,
	)
	assert.True(t, validation.Is(e))
	assert.StringContains(t, "Invalid body", e.Error())
}

func TestAnOpenAIServerErrorIsUnexpectedAndNotRetried(t *testing.T) {
	var calls atomic.Int64
	s, _ := upstream_tester.New(t, http.StatusOK, "")
	s.Config.Handler = http.HandlerFunc(
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			calls.Add(1)
			w.WriteHeader(http.StatusInternalServerError)
		},
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})
	assert.True(t, unexpected.Is(e))
	assert.Integer(t, 1, int(calls.Load()))
}

func TestABusyOpenAIServerIsUnavailable(t *testing.T) {
	e := openaiEmbedding(
		t,
		http.StatusServiceUnavailable,
		`{"error":{"message":"busy"}}`,
	)
	assert.True(t, unavailable.Is(e))
}
