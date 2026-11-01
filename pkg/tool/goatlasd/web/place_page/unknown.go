package place_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func Unknown() gomponents.Node {
	return html.P(
		html.Class(webConstant.AbsentClass),
		gomponents.Text(constant.PlaceUnknown),
	)
}
