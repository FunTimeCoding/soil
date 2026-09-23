package unit

import (
	"github.com/funtimecoding/soil/pkg/opnsense"
	"net/http/httptest"
)

func client(s *httptest.Server) *opnsense.Client {
	return opnsense.New(s.Listener.Addr().String(), "key", "secret", true)
}
