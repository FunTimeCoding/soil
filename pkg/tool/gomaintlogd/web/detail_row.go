package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
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
						htmx.Get(
							fragmentLocator(constant.EditPath, e.Identifier),
						),
						htmx.Target(target),
						htmx.Swap("outerHTML"),
						gomponents.Text("Edit"),
					),
					html.Button(
						html.Class("outline contrast"),
						htmx.Post(
							fragmentLocator(constant.DeletePath, e.Identifier),
						),
						htmx.Confirm("Delete this entry?"),
						gomponents.Attr(
							"hx-on::after-request",
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
