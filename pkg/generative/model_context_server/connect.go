package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"testing"
)

func (s *Server) Connect(t *testing.T) *model_context_client.Client {
	t.Helper()

	return model_context_client.New(t, s.Port, s.pool.Transport())
}
