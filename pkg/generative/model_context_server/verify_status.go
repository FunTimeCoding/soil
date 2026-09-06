package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func (s *Server) VerifyStatus(
	t *testing.T,
	path string,
	status int,
) {
	t.Helper()
	assert.HTTPStatus(t, s.address(path), status)
}
