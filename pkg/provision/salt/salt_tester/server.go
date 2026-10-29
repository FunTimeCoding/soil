package salt_tester

import (
	"net/http/httptest"
	"sync"
)

type Server struct {
	*httptest.Server
	mutex    sync.Mutex
	password string
	logins   int
	valid    string
	seen     []string
	answers  map[string]string
}
