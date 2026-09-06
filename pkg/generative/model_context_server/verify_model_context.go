package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"net/http"
	"testing"
)

func (s *Server) VerifyModelContext(t *testing.T) {
	t.Helper()
	assert.HTTPStatus(t, s.address("/mcp"), http.StatusUnauthorized)
	assert.HTTPStatus(t, s.address("/sse"), http.StatusUnauthorized)
	c := model_context_client.New(t, s.Port)
	defer c.Close()
	assert.True(t, len(c.ListTools()) > 0)
}
