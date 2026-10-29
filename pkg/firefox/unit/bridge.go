package unit

import (
	"github.com/funtimecoding/soil/pkg/firefox"
	"net/http/httptest"
	"testing"
)

func bridge(t *testing.T) (*firefox.Client, *httptest.Server) {
	t.Helper()
	c := firefox.New()
	s := httptest.NewServer(c)
	t.Cleanup(s.Close)
	t.Cleanup(c.Close)

	return c, s
}
