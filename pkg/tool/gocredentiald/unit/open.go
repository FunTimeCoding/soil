package unit

import (
	"github.com/funtimecoding/soil/pkg/tool/gocredentiald/service"
	"testing"
)

func open(
	t *testing.T,
	path string,
) *service.Service {
	t.Helper()

	return openRevealed(t, path, nil)
}
