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
	project := queryInteger(r, constant.Project)
	pipeline := queryInteger(r, constant.Pipeline)
	requested := queryInteger(r, constant.Job)
	jobs, e := s.client.PipelineJobs(project, pipeline)
	errors.PanicOnError(e)
	s.view.RenderFragment(
		w,
		s.logPane(project, pipeline, selectJob(jobs, requested)),
	)
}
