package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"testing"
)

func (s *Server) VerifyBase(t *testing.T) {
	t.Helper()
	assert.HTTPStatus(t, s.address("/health"), http.StatusOK)
	assert.HTTPStatus(t, s.address("/version"), http.StatusOK)
}
