package upstream_tester

import (
	"net/http/httptest"
	"sync"
)

type Tester struct {
	server  *httptest.Server
	lock    sync.Mutex
	logins  int
	tokens  []string
	written map[string]any
}
