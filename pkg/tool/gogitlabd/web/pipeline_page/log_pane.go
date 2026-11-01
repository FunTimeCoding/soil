package pipeline_page

import (
	"fmt"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/status"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/component"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) LogPane() gomponents.Node {
	j := p.view.Selected

	if j == nil {
		return html.Div(
			html.ID(constant.LogMark),
			html.P(
				html.Class(webConstant.AbsentClass),
				gomponents.Text("No jobs in this pipeline."),
			),
		)
	}

	attributes := []gomponents.Node{
		html.ID(constant.LogMark),
		html.Class(constant.LogClass),
	}

	if status.Active(j.Status) {
		attributes = append(
			attributes,
			extended.Get(
				fmt.Sprintf(
					"%s?%s&%s=%d",
					constant.JobPath,
					p.parameters(),
					argumentConstant.Job,
					j.Identifier,
				),
			),
			extended.Trigger("load delay:5s"),
			extended.Swap(webConstant.SwapOuter),
		)
	}

	header := []gomponents.Node{
		html.Class("log-header"),
		component.StatusIcon(j.Status),
		html.Strong(gomponents.Text(j.Name)),
		html.Small(gomponents.Text(j.Status)),
	}

	if j.Status == gitlabConstant.JobFail {
		header = append(
			header,
			action.New(
				fmt.Sprintf(
					"%s?%s&%s=%d",
					constant.RetryPath,
					p.parameters(),
					argumentConstant.Job,
					j.Identifier,
				),
				"Retry",
			).
				Swap(
					join.Empty("#", constant.DetailMark),
					webConstant.SwapOuter,
				).
				Class("retry").
				Small().
				Node(),
		)
	}

	body := gomponents.Text("No output yet.")

	if p.view.Trace != "" {
		body = gomponents.Text(formatTrace(p.view.Trace))
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
