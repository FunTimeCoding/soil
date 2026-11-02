package web

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/pipeline_page"
	"net/http"
)

func (s *Server) job(
	w http.ResponseWriter,
	r *http.Request,
) {
	instance, e := s.queryInstance(r)

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)

		return
	}

	v, f := s.service.JobView(
		r.Context(),
		instance,
		queryInteger(r, argumentConstant.Project),
		queryInteger(r, argumentConstant.Pipeline),
		queryInteger(r, argumentConstant.Job),
	)
	errors.PanicOnError(f)
	s.view.RenderFragment(w, pipeline_page.New(v).LogPane())
}
