package server

import (
	"github.com/funtimecoding/soil/pkg/web"
	"net/http"
	"net/http/pprof"
)

func (s *Server) Start() {
	s.Setup(s.Mux)

	if s.profiling {
		s.Mux.HandleFunc("GET /debug/pprof/", pprof.Index)
		s.Mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
		s.Mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
		s.Mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
		s.Mux.HandleFunc("GET /debug/pprof/trace", pprof.Trace)
	}

	var m http.Handler = s.Mux

	if s.Middleware != nil {
		m = s.Middleware(m)
	}

	s.http = web.Server(m, s.Address)

	if s.protected {
		s.http.ReadTimeout = readWriteTimeout
		s.http.WriteTimeout = readWriteTimeout

		if s.writeTimeout > 0 {
			s.http.WriteTimeout = s.writeTimeout
		}
	}

	if s.certificate != "" {
		web.ServeCertificateAsynchronous(
			s.http,
			s.certificate,
			s.key,
		)
	} else if s.listener != nil {
		web.ServeListenerAsynchronous(s.http, s.listener)
	} else {
		web.ServeAsynchronous(s.http)
	}
}
