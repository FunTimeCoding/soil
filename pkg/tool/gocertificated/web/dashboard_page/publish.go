package dashboard_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/change"
	"github.com/funtimecoding/soil/pkg/web/component/action"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Publish(changes []*change.Change) gomponents.Node {
	if len(changes) == 0 {
		return html.P(
			html.Class(webConstant.AbsentClass),
			gomponents.Text("Everything is published."),
		)
	}

	var items []gomponents.Node

	for _, c := range changes {
		items = append(items, html.Li(gomponents.Text(c.Path)))
	}

	return html.Div(
		html.H2(gomponents.Text(constant.PublishTitle)),
		html.Ul(gomponents.Group(items)),
		action.New(
			constant.PublishPath,
			fmt.Sprintf("Commit %d files", len(changes)),
		).Node(),
	)
}
