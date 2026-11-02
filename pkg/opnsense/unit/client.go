package unit

import (
	"github.com/funtimecoding/soil/pkg/opnsense"
	"net/http/httptest"
	"testing"
)

func client(
	t *testing.T,
	s *httptest.Server,
) *opnsense.Client {
	t.Helper()
	result := opnsense.New(s.Listener.Addr().String(), "key", "secret", true)
	t.Cleanup(result.Close)

	return result
}
