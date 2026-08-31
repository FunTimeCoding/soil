package web

import (
	"fmt"
	argument "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) pipelineDetail(
	project int64,
	pipeline int64,
	requested int64,
) gomponents.Node {
	jobs, e := s.client.PipelineJobs(project, pipeline)
	errors.PanicOnError(e)
	selected := selectJob(jobs, requested)
	p, f := s.client.Pipeline(project, pipeline)
	errors.PanicOnError(f)
	var ran gomponents.Node

	if p.CreatedAt != nil {
		ran = html.Small(
			gomponents.Text(p.CreatedAt.Local().Format(timeConstant.DateYear)),
		)
	}

	heading := gomponents.Textf("Pipeline #%d", pipeline)

	if len(jobs) > 0 && jobs[0].Project != nil {
		heading = gomponents.Textf(
			"%s #%d",
			join.Slash(
				[]string{jobs[0].Project.Namespace, jobs[0].Project.Name},
			),
			pipeline,
		)
	}

	external := html.Small(
		html.A(
			html.Href(p.WebURL),
			html.Target("_blank"),
			gomponents.Text("open in GitLab"),
		),
	)

	return html.Div(
		html.ID("pipeline-detail"),
		html.Div(
			html.Class("detail-header"),
			html.H3(heading),
			ran,
			external,
			html.Button(
				html.Class("delete"),
				gomponents.Attr(
					"hx-post",
					fmt.Sprintf(
						"%s?%s=%d&%s=%d",
						constant.DeletePath,
						argument.Project,
						project,
						argument.Pipeline,
						pipeline,
					),
				),
				gomponents.Attr(
					"hx-confirm",
					fmt.Sprintf("Delete pipeline #%d?", pipeline),
				),
				gomponents.Text("Delete"),
			),
		),
		html.Div(
			gomponents.Attr("sse-swap", constant.PipelineEvent),
			stageStrip(project, pipeline, jobs, selected),
		),
		s.logPane(project, pipeline, selected),
	)
}
