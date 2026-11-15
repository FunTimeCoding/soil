package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"net/http"
	"testing"
)

func TestOpenAIServerListingModelsReachable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(`{"object":"list","data":[]}`),
		upstream_server.WithObject(),
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.String(t, "/v1/models", s.Last().URL.Path)
}

func TestUnauthorizedOpenAIServerUnreachable(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithStatus(http.StatusUnauthorized),
		upstream_server.WithAnswer(`{"error":{"message":"bad key"}}`),
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	assert.NotNil(t, c.Reachable(t.Context()))
}
