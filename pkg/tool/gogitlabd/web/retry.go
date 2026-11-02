package web

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
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

	project := queryInteger(r, argumentConstant.Project)
	pipeline := queryInteger(r, argumentConstant.Pipeline)
	requested := queryInteger(r, argumentConstant.Job)
	x := r.Context()
	replacement, e := s.service.Retry(x, instance, project, requested)
	errors.PanicOnError(e)
	s.view.RenderFragment(
		w,
		s.detailRegion(x, instance, project, pipeline, replacement.Identifier),
	)
}
