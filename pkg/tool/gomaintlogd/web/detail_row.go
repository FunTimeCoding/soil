package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/time"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents-htmx"
	"maragu.dev/gomponents/html"
)

func detailRow(e *entry.Entry) gomponents.Node {
	target := fmt.Sprintf("#detail-%d", e.ID)

	return html.Tr(
		html.ID(fmt.Sprintf("detail-%d", e.ID)),
		html.Class("detail-row"),
		html.Td(
			gomponents.Attr("colspan", "6"),
			html.Div(
				html.Class("detail-content"),
				html.Div(
					html.Class("grid"),
					html.Div(
						html.Strong(gomponents.Text("Timestamp: ")),
						gomponents.Text(
							time.FormatCompact(e.Timestamp),
						),
					),
					html.Div(
						html.Strong(gomponents.Text("Action: ")),
						gomponents.Text(e.Action),
					),
					html.Div(
						html.Strong(gomponents.Text("User: ")),
						gomponents.Text(e.User),
					),
				),
				html.Div(
					html.Class("grid"),
					html.Div(
						html.Strong(gomponents.Text("System: ")),
						gomponents.Text(e.System),
					),
					html.Div(
						html.Strong(gomponents.Text("Service: ")),
						gomponents.Text(e.Service),
					),
					html.Div(),
				),
				html.Div(
					html.Strong(gomponents.Text("Description:")),
					html.Pre(gomponents.Text(e.Description)),
				),
				html.Div(
					html.Class("detail-actions"),
					html.Button(
						html.Class("outline"),
						htmx.Get(fmt.Sprintf("/entry/edit?id=%d", e.ID)),
						htmx.Target(target),
						htmx.Swap("outerHTML"),
						gomponents.Text("Edit"),
					),
					html.Button(
						html.Class("outline contrast"),
						htmx.Post(fmt.Sprintf("/entry/delete?id=%d", e.ID)),
						htmx.Confirm("Delete this entry?"),
						gomponents.Attr(
							"hx-on::after-request",
							fmt.Sprintf(
								"document.getElementById('row-%d')?.remove();document.getElementById('detail-%d')?.remove()",
								e.ID,
								e.ID,
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
								e.ID,
							),
						),
						gomponents.Text("Close"),
					),
				),
			),
		),
	)
}
