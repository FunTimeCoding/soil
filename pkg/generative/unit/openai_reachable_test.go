package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings/join"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAnOpenAIServerThatListsItsModelsIsReachable(t *testing.T) {
	var path string
	s := httptest.NewServer(
		http.HandlerFunc(
			func(
				w http.ResponseWriter,
				r *http.Request,
			) {
				path = r.URL.Path
				w.Header().Set(webConstant.ContentType, webConstant.Object)
				_, e := w.Write([]byte(`{"object":"list","data":[]}`))
				errors.PanicOnError(e)
			},
		),
	)
	defer s.Close()
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	assert.FatalOnError(t, c.Reachable(t.Context()))
	assert.String(t, "/v1/models", path)
}

func TestAnOpenAIServerThatRefusesIsNotReachable(t *testing.T) {
	s, _ := upstream_tester.New(
		t,
		http.StatusUnauthorized,
		`{"error":{"message":"bad key"}}`,
	)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	assert.NotNil(t, c.Reachable(t.Context()))
}
