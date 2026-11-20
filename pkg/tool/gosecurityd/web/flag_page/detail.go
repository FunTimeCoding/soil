package flag_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/overview_result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Detail(r *overview_result.Row) gomponents.Node {
	mitigation := MitigateForm(r.Flag.Vulnerability)

	if r.Flag.Mitigated {
		mitigation = html.P(
			html.Strong(gomponents.Text("Mitigated: ")),
			gomponents.Text(r.Flag.Mitigation),
		)
	}

	return opening(r).Detail(
		constant.FlagColumns,
		html.P(
			html.Strong(gomponents.Text("Note: ")),
			gomponents.Text(r.Flag.Note),
		),
		versions(r),
		html.P(
			html.A(
				html.Href(fragment(constant.LookupPath, r.Flag.Vulnerability)),
				gomponents.Text("Full lookup"),
			),
		),
		mitigation,
	).Response()
}
