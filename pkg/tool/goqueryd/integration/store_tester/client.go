package store_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/generative/ollama/reachable_skip"
	"testing"
)

func client(t *testing.T) *ollama.Client {
	t.Helper()
	result := ollama.NewEnvironment()
	t.Cleanup(result.Close)
	reachable_skip.Skip(t)

	return result
}
