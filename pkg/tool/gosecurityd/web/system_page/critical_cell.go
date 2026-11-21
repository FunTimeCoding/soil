package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func criticalCell(v *system_result.Summary) gomponents.Node {
	if v.Coverage == constant.CoverageUntracked {
		return html.Td(html.Small(gomponents.Text("unknown")))
	}

	return html.Td(CriticalBadge(v.Critical))
}
