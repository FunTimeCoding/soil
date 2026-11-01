package document_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store/result"
	"github.com/funtimecoding/soil/pkg/web/component/fact"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(d *result.Document) gomponents.Node {
	return gomponents.Group(
		[]gomponents.Node{
			html.H3(gomponents.Text(d.Title)),
			fact.New().
				Row("Collection", d.Collection).
				Row("Path", d.Path).
				RowNode(
					"Virtual Path",
					html.Small(gomponents.Text(d.VirtualPath)),
				).
				RowNode("File", html.Small(gomponents.Text(d.FilePath))).
				RowNode("Hash", html.Small(gomponents.Text(shortHash(d.Hash)))).
				Node(),
			html.H4(gomponents.Text("Body")),
			html.Pre(gomponents.Text(d.Body)),
		},
	)
}
