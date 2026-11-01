package component

import (
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/badge"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"time"
)

func StatusBadge(end *time.Time) gomponents.Node {
	if end == nil {
		return badge.New(constant.Firing).Color(webConstant.BadColor).Node()
	}

	return badge.New(constant.Resolved).Color(webConstant.OkayColor).Node()
}
