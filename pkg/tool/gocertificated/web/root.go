package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) root(
	w http.ResponseWriter,
	r *http.Request,
) {
	result, e := s.service.Authority(constant.RootAuthority)

	if e != nil {
		s.view.Answer(w, r, s.reporter, e)

		return
	}

	w.Header().Set(webConstant.ContentType, webConstant.Text)
	_, f := w.Write([]byte(result.Certificate))
	errors.PanicOnError(f)
}
