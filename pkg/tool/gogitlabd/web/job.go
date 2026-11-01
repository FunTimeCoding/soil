package web

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
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
		instance,
		queryInteger(r, constant.Project),
		queryInteger(r, constant.Pipeline),
		queryInteger(r, constant.Job),
	)
	errors.PanicOnError(f)
	s.view.RenderFragment(w, pipeline_page.New(v).LogPane())
}
