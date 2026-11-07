package model_context_server

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
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"testing"
)

func serve(
	t *testing.T,
	w *option.Web,
	setup func(*guard.Mux),
) *Server {
	t.Helper()
	p, n := system.ClaimPort()
	b := lifecycle.New(logger.New(context.Background())).WithServer(
		server.New(identity.Example(), w, memory.New(), setup).WithListener(n),
	)
	b.Run()
	assert.Listen(t, p)

	return &Server{Port: p, Lifecycle: b, pool: stall_client.New().Build()}
}
