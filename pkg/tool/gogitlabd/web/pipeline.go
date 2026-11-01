package web

import (
	"fmt"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/pipeline_page"
	"github.com/funtimecoding/soil/pkg/web/subscription"
	"net/http"
)

func (s *Server) pipeline(
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
	detail := s.detailRegion(instance, project, pipeline, requested)

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
				pipeline_page.Query(instance, project, pipeline),
			},
		),
		detail,
	)
}
