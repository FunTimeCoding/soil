package session_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"maragu.dev/gomponents"
	"time"
)

func RosterCruise(
	identifier string,
	setting *cruise_setting.Setting,
	grace time.Duration,
	alert string,
) gomponents.Node {
	mark := Event(identifier, constant.RosterCruisePart)

	return cruiseEntry(identifier, setting, grace).
		Identifier(mark).
		Swap(join.Empty("#", mark), webConstant.SwapOuter).
		Alert(alert).
		Node()
}
