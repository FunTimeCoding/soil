package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"testing"
)

func served(
	t *testing.T,
	path string,
	r *memory.Memory,
) *model_context_client.Client {
	t.Helper()

	return mounted(t, path, r).Connect(t)
}
