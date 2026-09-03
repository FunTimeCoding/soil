package model_context_server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	setup func(*http.ServeMux, *guard.Mux),
) *Server {
	t.Helper()
	p, n := system.ClaimPort()
	b := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New(
				identity.Example(),
				"",
				func(m *http.ServeMux) {
					setup(
						m,
						guard.New(m, []string{constant.ModelContextTestToken}),
					)
				},
			).WithListener(n),
		),
	)
	b.Run()
	assert.Listen(t, p)

	return &Server{Port: p, Lifecycle: b}
}
