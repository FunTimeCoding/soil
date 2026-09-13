package web

import (
	"fmt"
	argument "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (s *Server) logPane(
	project int64,
	pipeline int64,
	j *job.Job,
) gomponents.Node {
	if j == nil {
		return html.Div(
			html.ID("log-pane"),
			html.P(gomponents.Text("No jobs in this pipeline.")),
		)
	}

	trace, e := s.client.Trace(project, j.Identifier)
	errors.PanicOnError(e)
	attributes := []gomponents.Node{html.ID("log-pane"), html.Class("log-pane")}

	if board_entry.Active(j.Status) {
		attributes = append(
			attributes,
			extended.Get(
				fmt.Sprintf(
					"%s?%s=%d&%s=%d&%s=%d",
					constant.JobPath,
					argument.Project,
					project,
					argument.Pipeline,
					pipeline,
					argument.Job,
					j.Identifier,
				),
			),
			extended.Trigger("load delay:5s"),
			extended.Swap("outerHTML"),
		)
	}

	header := []gomponents.Node{
		html.Class("log-header"),
		html.Img(
			html.Class("status-icon"),
			html.Src(statusIcon(j.Status)),
			html.Alt(j.Status),
			html.Title(j.Status),
		),
		html.Strong(gomponents.Text(j.Name)),
		html.Small(gomponents.Text(j.Status)),
	}

	if j.Status == gitlabConstant.JobFail {
		header = append(
			header,
			html.Button(
				html.Class("retry"),
				extended.Post(
					fmt.Sprintf(
						"%s?%s=%d&%s=%d&%s=%d",
						constant.RetryPath,
						argument.Project,
						project,
						argument.Pipeline,
						pipeline,
						argument.Job,
						j.Identifier,
					),
				),
				extended.Target("#pipeline-detail"),
				extended.Swap("outerHTML"),
				gomponents.Text("Retry"),
			),
		)
	}

	body := gomponents.Text("No output yet.")

	if trace != "" {
		body = gomponents.Text(formatTrace(trace))
	}

	return html.Div(
		append(
			attributes,
			html.Div(header...),
			html.Pre(html.Code(body)),
			html.Script(gomponents.Raw(constant.ScrollScript)),
		)...,
	)
}
