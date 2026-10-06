package unit

import (
	"net/http/httptest"
	"sync"
)

type fakeIssuer struct {
	*httptest.Server
	mutex     sync.Mutex
	discovery int
}
