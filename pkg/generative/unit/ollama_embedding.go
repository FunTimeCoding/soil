package unit

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func ollamaEmbedding(
	t *testing.T,
	status int,
	answer string,
) error {
	t.Helper()
	s, _ := upstream_tester.New(t, status, answer)
	host, port := upstream_tester.HostPort(t, s.URL)
	c := ollama.New(ollama.WithHost(host), ollama.WithPort(port))
	defer c.Close()
	_, e := c.Embed(t.Context(), []string{"alfa"})

	return e
}
