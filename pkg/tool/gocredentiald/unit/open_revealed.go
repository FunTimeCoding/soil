package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"testing"
)

func openRevealed(
	t *testing.T,
	path string,
	revealedField []string,
) *service.Service {
	t.Helper()

	return openReported(t, path, revealedField, memory.New())
}
