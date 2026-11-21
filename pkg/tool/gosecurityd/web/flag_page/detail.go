package flag_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/affected_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(
	r *overview_result.Row,
	c *affected_count.Count,
) gomponents.Node {
	mitigation := MitigateForm(r.Flag.Subject)

	if r.Flag.Mitigated {
		mitigation = html.P(
			html.Strong(gomponents.Text("Mitigated: ")),
			gomponents.Text(r.Flag.Mitigation),
		)
	}

	note := html.P(
		html.Strong(gomponents.Text("Note: ")),
		gomponents.Text(r.Flag.Note),
	)

	if r.Debian == nil && r.Unread == nil {
		return opening(r, c).Detail(
			constant.FlagColumns,
			note,
			html.P(
				html.A(
					html.Href(constant.RemediationPath),
					gomponents.Text(constant.RemediationTitle),
				),
			),
			mitigation,
		).Response()
	}

	lookup := fragment(constant.LookupPath, r.Flag.Subject)

	return opening(r, c).Detail(
		constant.FlagColumns,
		note,
		DebianPips(r),
		versions(r),
		html.P(
			html.Strong(gomponents.Text("Affected systems: ")),
			gomponents.Text(hitSummary(c)),
			gomponents.Text(" "),
			html.A(
				html.Href(join.Empty(lookup, "#", constant.SystemsMark)),
				gomponents.Text("Which ones"),
			),
		),
		html.P(html.A(html.Href(lookup), gomponents.Text("Full lookup"))),
		mitigation,
	).Response()
}
