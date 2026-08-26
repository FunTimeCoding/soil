package server

import (
	"github.com/funtimecoding/soil/pkg/lifecycle/constant"
	"github.com/funtimecoding/soil/pkg/web"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/route"
	"net/http"
	"net/http/pprof"
)

func (s *Server) Start() {
	if s.identity.Stamp() == nil {
		panic(
			"identity has no stamp - argument.Parse stamps it, tests use identity.Example",
		)
	}

	s.Setup(s.Mux)
	s.Mux.HandleFunc(
		route.Get(webConstant.HealthPath),
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			w.WriteHeader(http.StatusOK)
		},
	)
	s.Mux.HandleFunc(
		route.Get(webConstant.VersionPath),
		func(
			w http.ResponseWriter,
			_ *http.Request,
		) {
			b := s.identity.Stamp()
			web.EncodeNotation(
				w,
				&Version{
					Name:      s.identity.Name(),
					Version:   b.Version,
					GitHash:   b.GitHash,
					BuildDate: b.BuildDate,
				},
			)
		},
	)

	if s.profiling {
		s.Mux.HandleFunc("GET /debug/pprof/", pprof.Index)
		s.Mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
		s.Mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
		s.Mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
		s.Mux.HandleFunc("GET /debug/pprof/trace", pprof.Trace)
	}

	var m http.Handler = s.Mux

	if len(s.tokens) > 0 {
		m = web.TokenMiddleware(
			s.tokens,
			webConstant.HealthPath,
			webConstant.VersionPath,
		)(m)
	}

	if s.Middleware != nil {
		m = s.Middleware(m)
	}

	s.http = web.Server(m, s.Address)

	if s.protected {
		s.http.ReadTimeout = constant.ReadWriteTimeout
		s.http.WriteTimeout = constant.ReadWriteTimeout

		if s.writeTimeout > 0 {
			s.http.WriteTimeout = s.writeTimeout
		}
	}

	if s.certificate != "" {
		web.ServeCertificateAsynchronous(s.http, s.certificate, s.key)
	} else if s.listener != nil {
		web.ServeListenerAsynchronous(s.http, s.listener)
	} else {
		web.ServeAsynchronous(s.http)
	}
}
