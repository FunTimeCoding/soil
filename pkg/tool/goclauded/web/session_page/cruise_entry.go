package session_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"github.com/funtimecoding/soil/pkg/web/component/choice"
	"github.com/funtimecoding/soil/pkg/web/component/entry"
	"strconv"
	"time"
)

func cruiseEntry(
	identifier string,
	setting *cruise_setting.Setting,
	grace time.Duration,
) *entry.Entry {
	kind := constant.CruiseMinutes
	minutes := ""

	switch {
	case setting.Grace == nil:
		kind = constant.CruiseDefault
	case *setting.Grace == 0:
		kind = constant.CruiseOff
	}

	if grace > 0 {
		minutes = strconv.Itoa(int(grace.Minutes()))
	}

	kinds := []*choice.Choice{
		choice.New(constant.CruiseDefault, constant.CruiseDefault),
		choice.New(constant.CruiseOff, constant.CruiseOff),
		choice.New(constant.CruiseMinutes, constant.CruiseMinutes),
	}

	return entry.New(ActionPath(identifier, constant.CruiseAction), "Set").
		Inline().
		Select(constant.GraceKind, "Cruise", kinds, kind).
		Number(constant.Grace, "Minutes", minutes).
		Min(strconv.Itoa(constant.CruiseGraceMinimumMinutes)).
		Max(strconv.Itoa(constant.CruiseGraceMaximumMinutes)).
		Step("1")
}
