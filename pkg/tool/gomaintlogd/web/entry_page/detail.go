package entry_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/web/component/load"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(e *entry.Entry) gomponents.Node {
	mark := detailMark(e.Identifier)
	target := join.Empty("#", mark)

	return html.Tr(
		html.ID(mark),
		html.Class(constant.DetailRowClass),
		html.Td(
			gomponents.Attr("colspan", "6"),
			html.Div(
				html.Class(constant.DetailContentClass),
				fields(e),
				html.Div(
					html.Class(constant.DetailActionsClass),
					load.New(fragment(constant.EditPath, e.Identifier), "Edit").
						Swap(target, webConstant.SwapOuter).
						Button().
						Class("outline").
						Node(),
					html.Button(
						html.Class("outline contrast"),
						extended.Post(
							fragment(constant.DeletePath, e.Identifier),
						),
						extended.Confirm("Delete this entry?"),
						extended.Swap(webConstant.SwapNone),
						extended.AfterRequest(
							fmt.Sprintf(
								"document.getElementById('%s%d')?.remove();document.getElementById('%s')?.remove()",
								constant.RowPrefix,
								e.Identifier,
								mark,
							),
						),
						gomponents.Text("Delete"),
					),
					html.Button(
						html.Type("button"),
						html.Class("outline secondary"),
						gomponents.Attr(
							"onclick",
							fmt.Sprintf(
								"var r=document.getElementById('%s');r.hidden=true;r.innerHTML='';r.className=''",
								mark,
							),
						),
						gomponents.Text("Close"),
					),
					load.New(Locator(e.Identifier), "Permalink").
						Button().
						Class("outline secondary").
						Node(),
				),
			),
		),
	)
}
