package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
	"testing"
)

func serveTrace(
	t *testing.T,
	setup func(*http.ServeMux),
) int {
	t.Helper()
	p, n := system.ClaimPort()
	l := lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(identity.Example(), "", setup).WithListener(n),
	)
	l.Run()
	t.Cleanup(l.Close)
	assert.Listen(t, p)

	return p
}
