package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"os"
)

func reportsTable(v []os.DirEntry) gomponents.Node {
	if len(v) == 0 {
		return html.P(html.Em(gomponents.Text("No reports generated yet.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Report")),
				html.Th(gomponents.Text("Size")),
				html.Th(),
			),
		),
		html.TBody(
			gomponents.Map(
				v,
				func(d os.DirEntry) gomponents.Node {
					i, e := d.Info()
					errors.PanicOnError(e)

					return html.Tr(
						html.Td(
							html.A(
								html.Href(
									fmt.Sprintf(
										"/reports/%s",
										d.Name(),
									),
								),
								gomponents.Text(d.Name()),
							),
						),
						html.Td(
							html.Class("report-size"),
							gomponents.Text(formatSize(i.Size())),
						),
						html.Td(
							html.Form(
								html.Method("post"),
								html.Action(
									fmt.Sprintf(
										"/reports/%s/delete",
										d.Name(),
									),
								),
								html.Button(
									html.Type("submit"),
									html.Class("secondary outline"),
									gomponents.Text("Delete"),
								),
							),
						),
					)
				},
			),
		),
	)
}
