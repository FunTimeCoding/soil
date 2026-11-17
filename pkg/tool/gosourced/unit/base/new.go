package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{
		Memory:         mock_client.New(t),
		Reporter:       memory.New(),
		inventory:      inventory.New(),
		indexDirectory: t.TempDir(),
		t:              t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
