package issuer_server

import (
	"net/http/httptest"
	"sync"
	"testing"
)

type Server struct {
	*httptest.Server
	t         *testing.T
	mutex     sync.Mutex
	discovery int
}
