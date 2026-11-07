package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
	"testing"
)

func serveTrace(
	t *testing.T,
	setup func(*guard.Mux),
) int {
	t.Helper()
	p, n := system.ClaimPort()
	l := lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(identity.Example(), option.New(""), memory.New(), setup).
			WithListener(n),
	)
	l.Run()
	t.Cleanup(l.Close)
	assert.Listen(t, p)

	return p
}
