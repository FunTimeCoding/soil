package upstream_server

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

type Server struct {
	*httptest.Server
	t             *testing.T
	mutex         sync.Mutex
	statuses      []int
	answer        string
	header        http.Header
	authorization string
	requests      []*http.Request
	opened        int
	secure        bool
}
