package upgrade_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/web/component/opening_row"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"strconv"
)

func opening(g *upgrade_group.Group) *opening_row.OpeningRow {
	return opening_row.New(
		identifier(g),
		fragment(constant.UpgradeOpenPath, g),
		fragment(constant.UpgradeShutPath, g),
	).Cells(
		html.Td(gomponents.Text(g.Source)),
		html.Td(gomponents.Text(g.Release)),
		html.Td(system_page.ScoreBadge(g.TopScore)),
		html.Td(gomponents.Text(strconv.Itoa(len(g.Hosts)))),
		html.Td(gomponents.Text(strconv.Itoa(len(g.Vulnerabilities)))),
		html.Td(html.Small(gomponents.Text(installed(g)))),
		html.Td(html.Small(gomponents.Text(g.Candidate))),
	)
}
