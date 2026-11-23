package graph_server

import (
	"net/http/httptest"
	"sync"
)

type Server struct {
	server    *httptest.Server
	lock      sync.Mutex
	answers   map[string][]Answer
	exchanges []Exchange
}
