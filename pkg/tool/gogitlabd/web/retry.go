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
	project := queryInteger(r, constant.Project)
	pipeline := queryInteger(r, constant.Pipeline)
	requested := queryInteger(r, constant.Job)
	replacement, e := s.client.Retry(project, requested)
	errors.PanicOnError(e)
	s.view.RenderFragment(
		w,
		s.pipelineDetail(project, pipeline, replacement.Identifier),
	)
}
