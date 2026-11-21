package system_page

import (
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"strconv"
)

func CriticalBadge(n int) gomponents.Node {
	if n == 0 {
		return gomponents.Text(strconv.Itoa(n))
	}

	return badge.New(strconv.Itoa(n)).Color(webConstant.BadColor).Node()
}
