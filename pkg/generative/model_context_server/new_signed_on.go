package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
	"testing"
)

func NewSignedOn(
	t *testing.T,
	authorization func() *client.Client,
	setup func(*guard.Mux),
) *Server {
	t.Helper()
	w := option.New("", constant.ModelContextTestToken)
	w.Authorization = authorization

	return serve(t, w, setup)
}
