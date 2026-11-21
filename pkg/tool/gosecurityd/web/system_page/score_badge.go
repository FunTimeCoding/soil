package system_page

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func ScoreBadge(score float64) gomponents.Node {
	if score == 0 {
		return html.Small(gomponents.Text("not scored"))
	}

	color := webConstant.QuietColor

	switch {
	case score >= constant.CriticalScore:
		color = webConstant.BadColor
	case score >= constant.HighScore:
		color = webConstant.WarnColor
	}

	return badge.New(fmt.Sprintf("%.1f", score)).Color(color).Node()
}
