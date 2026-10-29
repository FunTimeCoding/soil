package salt_tester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	password string,
) *Server {
	t.Helper()
	result := &Server{password: password, answers: map[string]string{}}
	result.Server = httptest.NewServer(http.HandlerFunc(result.serve))
	t.Cleanup(result.Close)

	return result
}
