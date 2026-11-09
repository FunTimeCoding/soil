package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"strconv"
)

func cruiseEntry(
	identifier string,
	setting *cruise_setting.Setting,
) *entry.Entry {
	pace := ""

	if setting.Mode == constant.CruisePaced {
		pace = strconv.Itoa(int(setting.Pace.Minutes()))
	}

	modes := []*choice.Choice{
		choice.New(string(constant.CruiseOff), string(constant.CruiseOff)),
		choice.New(string(constant.CruiseOn), string(constant.CruiseOn)),
		choice.New(string(constant.CruisePaced), string(constant.CruisePaced)),
	}

	return entry.New(ActionPath(identifier, constant.CruiseAction), "Set").
		Inline().
		Select(constant.Mode, "Cruise", modes, string(setting.Mode)).
		Number(constant.Pace, "Minutes", pace).
		Min(strconv.Itoa(constant.CruisePaceMinimumMinutes)).
		Max(strconv.Itoa(constant.CruisePaceMaximumMinutes)).
		Step("1")
}
