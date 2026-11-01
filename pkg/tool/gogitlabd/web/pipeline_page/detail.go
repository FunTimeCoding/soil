package pipeline_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	timeConstant "github.com/funtimecoding/soil/pkg/time/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func (p *Page) Detail() gomponents.Node {
	v := p.view
	var ran gomponents.Node

	if v.Detail.Create != nil {
		ran = html.Small(
			gomponents.Text(
				v.Detail.Create.Local().Format(timeConstant.DateYear),
			),
		)
	}

	heading := gomponents.Textf("Pipeline #%d", v.Pipeline)

	if len(v.Jobs) > 0 && v.Jobs[0].Project != nil {
		heading = gomponents.Textf(
			"%s #%d",
			join.Slash(
				[]string{v.Jobs[0].Project.Namespace, v.Jobs[0].Project.Name},
			),
			v.Pipeline,
		)
	}

	return html.Div(
		html.ID(constant.DetailMark),
		html.Div(
			html.Class("detail-header"),
			html.H3(heading),
			ran,
			html.Small(
				html.A(
					html.Href(v.Detail.Link),
					html.Target("_blank"),
					gomponents.Text("open in GitLab"),
				),
			),
			action.New(
				fmt.Sprintf("%s?%s", constant.DeletePath, p.parameters()),
				"Delete",
			).
				Swap(
					join.Empty("#", constant.DetailMark),
					webConstant.SwapOuter,
				).
				Confirm(fmt.Sprintf("Delete pipeline #%d?", v.Pipeline)).
				Class("delete").
				Small().
				Node(),
		),
		html.Div(extended.StreamSwap(constant.PipelineEvent), p.StageStrip()),
		p.LogPane(),
	)
}
