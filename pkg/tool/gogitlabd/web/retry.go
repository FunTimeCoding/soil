package web

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (s *Server) retry(
	w http.ResponseWriter,
	r *http.Request,
) {
	instance, e := s.queryInstance(r)

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)

		return
	}

	project := queryInteger(r, constant.Project)
	pipeline := queryInteger(r, constant.Pipeline)
	requested := queryInteger(r, constant.Job)
	replacement, e := s.service.Retry(instance, project, requested)
	errors.PanicOnError(e)
	s.view.RenderFragment(
		w,
		s.detailRegion(instance, project, pipeline, replacement.Identifier),
	)
}
