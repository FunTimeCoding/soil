package server

import (
	"github.com/funtimecoding/soil/pkg/identity"
	"net"
	"net/http"
	"time"
)

type Server struct {
	Mux          *http.ServeMux
	http         *http.Server
	Setup        func(*http.ServeMux)
	Middleware   func(http.Handler) http.Handler
	identity     *identity.Tool
	Address      string
	listener     net.Listener
	tokens       []string
	protected    bool
	writeTimeout time.Duration
	certificate  string
	key          string
	profiling    bool
}
