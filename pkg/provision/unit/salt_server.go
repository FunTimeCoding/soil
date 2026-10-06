package unit

import (
	"net/http/httptest"
	"sync"
)

type saltServer struct {
	*httptest.Server
	mutex  sync.Mutex
	logins int
	valid  string
	seen   []string
}
