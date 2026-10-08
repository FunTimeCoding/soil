package unit

import (
	"net/http/httptest"
	"sync"
)

type FakeIssuer struct {
	*httptest.Server
	mutex     sync.Mutex
	discovery int
}
