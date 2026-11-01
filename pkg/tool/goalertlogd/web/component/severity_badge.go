package component

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func SeverityBadge(severity string) gomponents.Node {
	color, known := constant.SeverityColor[severity]

	if !known {
		color = webConstant.QuietColor
	}

	return badge.New(severity).Color(color).Node()
}
