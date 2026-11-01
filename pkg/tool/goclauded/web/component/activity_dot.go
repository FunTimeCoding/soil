package component

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/web/component/status_dot"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"time"
)

func ActivityDot(lastSeen time.Time) gomponents.Node {
	if time.Since(lastSeen) > constant.RosterStaleAfter {
		return status_dot.New(webConstant.QuietColor).
			Title(constant.IdleLabel).
			Node()
	}

	return status_dot.New(webConstant.OkayColor).
		Title(constant.ActiveLabel).
		Node()
}
