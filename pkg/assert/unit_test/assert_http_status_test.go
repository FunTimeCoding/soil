package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPStatus(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, q *http.Request) {
				w.WriteHeader(http.StatusTeapot)
			},
		),
	)
	defer server.Close()
	assert.HTTPStatus(t, server.URL, http.StatusTeapot)
}
