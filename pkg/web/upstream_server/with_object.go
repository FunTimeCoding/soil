package upstream_server

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func WithObject() Option {
	return func(s *Server) {
		if s.header == nil {
			s.header = http.Header{}
		}

		s.header.Set(constant.ContentType, constant.Object)
	}
}
