package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/option"
	"testing"
)

func New(
	t *testing.T,
	setup func(*guard.Mux),
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{}

	for _, f := range o {
		f(result)
	}

	w := option.New("", constant.ModelContextTestToken)

	if result.signOn {
		w.Authorization = authorization_tester.Client
	}

	serve(t, result, w, setup)
	t.Cleanup(result.Close)

	return result
}
