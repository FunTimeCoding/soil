package openbao_server

import (
	"net/http/httptest"
	"sync"
)

type Server struct {
	server  *httptest.Server
	lock    sync.Mutex
	logins  int
	tokens  []string
	written map[string]any
}
