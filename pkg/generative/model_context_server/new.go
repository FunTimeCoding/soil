package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
	"testing"
)

func New(
	t *testing.T,
	setup func(*guard.Mux),
) *Server {
	t.Helper()

	return serve(t, option.New("", constant.ModelContextTestToken), setup)
}
