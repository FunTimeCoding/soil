package roster_card

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/pulse"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

type Card struct {
	Session *session.Session
	Labels  []label.Label
	Latest  *pulse.Pulse
	Queued  int64
	Cruise  *cruise_setting.Setting
}
