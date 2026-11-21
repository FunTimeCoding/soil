package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func KernelBadge(state string) gomponents.Node {
	color := webConstant.WarnColor

	switch state {
	case constant.KernelCurrent:
		color = webConstant.OkayColor
	case constant.KernelStuck:
		color = webConstant.BadColor
	case constant.KernelUnknown:
		color = webConstant.QuietColor
	}

	return badge.New(state).Color(color).Node()
}
