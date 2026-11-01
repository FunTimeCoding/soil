package alert_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"time"
)

func Filter(
	start time.Time,
	end time.Time,
) gomponents.Node {
	return entry.New(constant.RecentPath, "Filter").
		Query().
		Swap(join.Empty("#", constant.RecentMark), webConstant.SwapInner).
		Inline().
		DateTime(constant.Start, "Start", start.Format(constant.FilterFormat)).
		DateTime(constant.End, "End", end.Format(constant.FilterFormat)).
		Node()
}
