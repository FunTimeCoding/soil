package flag_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"github.com/funtimecoding/soil/pkg/web/component/status_dot"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func DebianPips(r *overview_result.Row) gomponents.Node {
	label := html.Strong(gomponents.Text("Debian: "))

	if r.Unread != nil {
		return html.P(label, html.Small(gomponents.Text(r.Unread.Error())))
	}

	if len(r.Debian.Statuses) == 0 {
		return html.P(label, html.Small(gomponents.Text("no entries")))
	}

	pips := make([]gomponents.Node, 0, len(r.Debian.Statuses))

	for _, s := range r.Debian.Statuses {
		color := webConstant.BadColor
		title := fmt.Sprintf("%s open", s.Package)

		switch {
		case s.NeverAffected:
			color = webConstant.QuietColor
			title = fmt.Sprintf("%s not affected", s.Package)
		case s.Fixed != "":
			color = webConstant.OkayColor
			title = fmt.Sprintf("%s fixed in %s", s.Package, s.Fixed)
		}

		pips = append(
			pips,
			html.Span(
				html.Title(title),
				status_dot.New(color).Node(),
				gomponents.Text(" "),
				html.Small(gomponents.Text(s.Release)),
			),
		)
	}

	return html.P(
		label,
		html.Span(html.Class(webConstant.ClusterClass), gomponents.Group(pips)),
	)
}
