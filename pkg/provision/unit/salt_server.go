package unit

import (
	"net/http/httptest"
	"sync"
)

type SaltServer struct {
	*httptest.Server
	mutex  sync.Mutex
	logins int
	valid  string
	seen   []string
}
