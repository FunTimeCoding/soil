package remediation_page

import (
	"fmt"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/system_page"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/web/upgrade_page"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/url"
)

func itemCell(i *remediation_item.Item) gomponents.Node {
	label := i.Subject
	link := system_page.Query(
		constant.LookupPath,
		url.Values{generativeConstant.ParameterIdentifier: {i.Subject}},
	)

	switch i.Kind {
	case constant.RemediationUpgrade:
		label = fmt.Sprintf("%s on %s", i.Source, i.Release)
		link = system_page.Query(
			constant.UpgradesPath,
			upgrade_page.Kept(i.Release, i.Kernel),
		)
	case constant.RemediationRelease,
		constant.RemediationReboot,
		constant.RemediationKernel:
		label = fmt.Sprintf("%s (%s)", i.Hosts[0], i.Release)
		link = system_page.Query(
			constant.SystemPath,
			url.Values{generativeConstant.ParameterName: {i.Hosts[0]}},
		)
	}

	return html.Td(
		html.A(html.Href(link), gomponents.Text(label)),
		html.Br(),
		html.Small(html.Code(gomponents.Text(i.Subject))),
	)
}
