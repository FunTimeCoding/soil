package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func stageStrip(
	project int64,
	pipeline int64,
	jobs []*job.Job,
	selected *job.Job,
) gomponents.Node {
	order := []string{}
	grouped := map[string][]*job.Job{}

	for _, j := range jobs {
		if _, okay := grouped[j.Stage]; !okay {
			order = append(order, j.Stage)
		}

		grouped[j.Stage] = append(grouped[j.Stage], j)
	}

	groups := []gomponents.Node{html.Class("stage-strip")}

	for _, stage := range order {
		chips := []gomponents.Node{
			html.Class("stage-group"),
			html.Span(html.Class("stage-name"), gomponents.Text(stage)),
		}

		for _, j := range grouped[stage] {
			class := "job-chip"

			if selected != nil && j.Identifier == selected.Identifier {
				class = "job-chip selected"
			}

			chips = append(
				chips,
				html.A(
					html.Class(class),
					gomponents.Attr(
						"hx-get",
						detailLink(project, pipeline, j.Identifier),
					),
					gomponents.Attr("hx-target", "#pipeline-detail"),
					gomponents.Attr("hx-swap", "outerHTML"),
					html.Img(
						html.Class("status-icon"),
						html.Src(statusIcon(j.Status)),
						html.Alt(j.Status),
						html.Title(j.Status),
					),
					gomponents.Text(j.Name),
				),
			)
		}

		groups = append(groups, html.Div(chips...))
	}

	return html.Div(groups...)
}
