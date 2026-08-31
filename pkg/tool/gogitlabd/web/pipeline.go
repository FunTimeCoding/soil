package web

import (
	"fmt"
	argument "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
	"strconv"
)

func (s *Server) pipeline(
	w http.ResponseWriter,
	r *http.Request,
) {
	project := queryInteger(r, argument.Project)
	pipeline := queryInteger(r, argument.Pipeline)
	requested := queryInteger(r, argument.Job)
	detail := s.pipelineDetail(project, pipeline, requested)

	if s.view.IsExtendedRequest(r) {
		s.view.RenderFragment(w, detail)

		return
	}

	s.view.RenderLivePage(
		w,
		fmt.Sprintf("Pipeline #%d", pipeline),
		constant.PipelinePath,
		join.Ampersand(
			[]string{
				subscription.Query(constant.PipelineEvent),
				key_value.Equals(
					argument.Project,
					strconv.FormatInt(project, 10),
				),
				key_value.Equals(
					argument.Pipeline,
					strconv.FormatInt(pipeline, 10),
				),
			},
		),
		detail,
	)
}
