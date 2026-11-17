package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/unit/daemon_tester"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New(t)
	r := memory.New()
	result := &Server{
		MockClient: c,
		Reporter:   r,
		Server:     model_context_server.Mount(t, daemon_tester.New(t, c), r),
	}
	t.Cleanup(result.Close)

	return result
}
