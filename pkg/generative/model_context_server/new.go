package model_context_server

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	setup func(*http.ServeMux),
) *Server {
	t.Helper()
	p, n := system.ClaimPort()
	b := lifecycle.New(
		logger.New(context.Background()),
		lifecycle.WithServer(
			server.New("", setup).WithListener(n),
		),
	)
	b.Run()
	assert.Listen(t, p)

	return &Server{Port: p, Lifecycle: b}
}
