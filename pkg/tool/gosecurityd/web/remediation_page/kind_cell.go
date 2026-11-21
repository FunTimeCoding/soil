package remediation_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/remediation_item"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func kindCell(i *remediation_item.Item) gomponents.Node {
	text := map[string]string{
		constant.RemediationUpgrade: "upgrade",
		constant.RemediationRelease: "release upgrade",
		constant.RemediationReboot:  "reboot into the installed kernel",
		constant.RemediationKernel: join.Space(
			"kernel out of reach:",
			i.Reason,
		),
	}[i.Kind]

	if i.Count != nil {
		text = fmt.Sprintf(
			"no upgrade: Debian fix not in candidate %d, paid only %d, no fix %d",
			i.Count.Pending,
			i.Count.Paid,
			i.Count.Open,
		)
	}

	return html.Td(html.Small(gomponents.Text(text)))
}
