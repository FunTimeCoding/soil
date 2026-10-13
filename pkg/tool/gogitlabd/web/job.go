package web

import (
	"github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
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

	project := queryInteger(r, constant.Project)
	pipeline := queryInteger(r, constant.Pipeline)
	requested := queryInteger(r, constant.Job)
	jobs, e := s.service.PipelineJobs(instance, project, pipeline)
	errors.PanicOnError(e)
	s.view.RenderFragment(
		w,
		s.logPane(instance, project, pipeline, selectJob(jobs, requested)),
	)
}
