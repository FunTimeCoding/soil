package session_view

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/context_load"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/session_detail"
)

type View struct {
	Detail      *session_detail.Detail
	Labels      []label.Label
	Loads       []context_load.Load
	LoadsHidden string
	Pulses      []pulse.Pulse
	Cruise      *cruise_setting.Setting
	DeleteHash  string
}
