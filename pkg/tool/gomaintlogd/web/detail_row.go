package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/extended"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func detailRow(e *entry.Entry) gomponents.Node {
	target := fmt.Sprintf("#detail-%d", e.Identifier)

	return html.Tr(
		html.ID(fmt.Sprintf("detail-%d", e.Identifier)),
		html.Class("detail-row"),
		html.Td(
			gomponents.Attr("colspan", "6"),
			html.Div(
				html.Class("detail-content"),
				entryFields(e),
				html.Div(
					html.Class("detail-actions"),
					html.Button(
						html.Class("outline"),
						extended.Get(
							fragmentLocator(constant.EditPath, e.Identifier),
						),
						extended.Target(target),
						extended.Swap(webConstant.SwapOuter),
						gomponents.Text("Edit"),
					),
					html.Button(
						html.Class("outline contrast"),
						extended.Post(
							fragmentLocator(constant.DeletePath, e.Identifier),
						),
						extended.Confirm("Delete this entry?"),
						extended.Swap(webConstant.SwapNone),
						extended.AfterRequest(
							fmt.Sprintf(
								"document.getElementById('row-%d')?.remove();document.getElementById('detail-%d')?.remove()",
								e.Identifier,
								e.Identifier,
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
								"var r=document.getElementById('detail-%d');r.style.display='none';r.innerHTML='';r.className=''",
								e.Identifier,
							),
						),
						gomponents.Text("Close"),
					),
					html.A(
						html.Class("outline secondary"),
						gomponents.Attr("role", "button"),
						html.Href(entryLocator(e.Identifier)),
						gomponents.Text("Permalink"),
					),
				),
			),
		),
	)
}
