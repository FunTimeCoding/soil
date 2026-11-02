package upstream_tester

import (
	"net/http/httptest"
	"sync"
	"testing"
)

type Upstream struct {
	t       *testing.T
	server  *httptest.Server
	mutex   sync.Mutex
	answers map[string]*Answer
}
