package web

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) activityPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		s.activitySection(r.URL.Query()[constant.Kind]).Render(w),
	)
}
