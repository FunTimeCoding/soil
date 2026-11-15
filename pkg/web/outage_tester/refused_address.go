package outage_tester

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func RefusedAddress(t *testing.T) string {
	t.Helper()
	s := httptest.NewServer(http.NotFoundHandler())
	s.Close()

	return s.URL
}
