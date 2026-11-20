package system_page

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
)

func CoverageBadge(coverage string) gomponents.Node {
	color := webConstant.BadColor

	switch coverage {
	case constant.CoverageSupported:
		color = webConstant.OkayColor
	case constant.CoverageTracked:
		color = webConstant.WarnColor
	}

	return badge.New(coverage).Color(color).Node()
}
