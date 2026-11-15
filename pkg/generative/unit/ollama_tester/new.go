package ollama_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func New(
	t *testing.T,
	address string,
	o ...ollama.Option,
) *ollama.Client {
	t.Helper()
	host, port := upstream_server.HostPort(t, address)
	result := ollama.New(
		append(
			[]ollama.Option{ollama.WithHost(host), ollama.WithPort(port)},
			o...,
		)...,
	)
	t.Cleanup(result.Close)

	return result
}
