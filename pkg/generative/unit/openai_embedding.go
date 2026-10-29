package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/openai"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func openaiEmbedding(
	t *testing.T,
	status int,
	answer string,
) error {
	t.Helper()
	s, _ := upstream_tester.New(t, status, answer)
	c := openai.NewBase(join.Empty(s.URL, "/v1"), constant.OpenAIEmbedModel)
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})

	return e
}
