package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
	"net/http"
)

func (s *Server) root(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.store.Authority(constant.RootAuthority)
	errors.PanicOnError(e)

	if result == nil {
		http.Error(w, constant.RootMissing, http.StatusNotFound)

		return
	}

	w.Header().Set(web.ContentType, web.Text)
	_, f := w.Write([]byte(result.Certificate))
	errors.PanicOnError(f)
}
