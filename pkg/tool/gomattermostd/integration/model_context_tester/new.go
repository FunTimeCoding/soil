package model_context_tester

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/integration/base"
	"net/http"
	"testing"
)

func New(t *testing.T, configure func(*http.ServeMux)) *Tester {
	t.Helper()
	s := base.New(t, configure)
	c := model_context_client.New(t, s.ContextServer.Port)
	t.Cleanup(
		func() {
			c.Close()
			s.Close()
		},
	)

	return &Tester{Client: c, Server: s}
}
