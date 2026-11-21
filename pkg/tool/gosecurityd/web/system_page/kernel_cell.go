package system_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func kernelCell(v *system_result.Summary) gomponents.Node {
	detail := v.RunningKernel

	if v.KernelReason != "" {
		detail = join.CommaSpace([]string{v.RunningKernel, v.KernelReason})
	}

	return html.Td(
		KernelBadge(v.KernelState),
		html.Br(),
		html.Small(gomponents.Text(detail)),
	)
}
