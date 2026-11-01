package pipeline_page

import (
	"github.com/funtimecoding/soil/pkg/gitlab/job"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web/component"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) StageStrip() gomponents.Node {
	v := p.view
	var order []string
	grouped := map[string][]*job.Job{}

	for _, j := range v.Jobs {
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

			if v.Selected != nil && j.Identifier == v.Selected.Identifier {
				class = "job-chip selected"
			}

			chips = append(
				chips,
				html.A(
					html.Class(class),
					extended.Get(
						Path(v.Instance, v.Project, v.Pipeline, j.Identifier),
					),
					extended.Target(join.Empty("#", constant.DetailMark)),
					extended.Swap(webConstant.SwapOuter),
					component.StatusIcon(j.Status),
					gomponents.Text(j.Name),
				),
			)
		}

		groups = append(groups, html.Div(chips...))
	}

	return html.Div(groups...)
}
