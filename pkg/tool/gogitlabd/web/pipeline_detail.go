package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) pipelineDetail(
	instance string,
	project int64,
	pipeline int64,
	requested int64,
) gomponents.Node {
	jobs, e := s.service.PipelineJobs(instance, project, pipeline)
	errors.PanicOnError(e)
	selected := selectJob(jobs, requested)
	p, f := s.service.Pipeline(instance, project, pipeline)
	errors.PanicOnError(f)
	var ran gomponents.Node

	if p.Create != nil {
		ran = html.Small(
			gomponents.Text(p.Create.Local().Format(timeConstant.DateYear)),
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
			html.Href(p.Link),
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
				extended.Post(
					fmt.Sprintf(
						"%s?%s",
						constant.DeletePath,
						pipelineQuery(instance, project, pipeline),
					),
				),
				extended.Confirm(fmt.Sprintf("Delete pipeline #%d?", pipeline)),
				gomponents.Text("Delete"),
			),
		),
		html.Div(
			extended.StreamSwap(constant.PipelineEvent),
			stageStrip(instance, project, pipeline, jobs, selected),
		),
		s.logPane(instance, project, pipeline, selected),
	)
}
