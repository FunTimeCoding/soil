package server

import (
	"net"
	"net/http"
	"time"
)

type Server struct {
	Mux          *http.ServeMux
	http         *http.Server
	Setup        func(*http.ServeMux)
	Middleware   func(http.Handler) http.Handler
	Address      string
	listener     net.Listener
	protected    bool
	writeTimeout time.Duration
	certificate  string
	key          string
	profiling    bool
}
