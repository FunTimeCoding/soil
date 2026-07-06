package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/integration_test/base"
	"testing"
)

func New(
	t *testing.T,
	fixture string,
) *Tester {
	t.Helper()
	s := base.New(t, fixture)

	return &Tester{
		server:    s,
		Client:    model_context_client.New(t, s.ContextServer.Port),
		Directory: s.Directory,
	}
}
