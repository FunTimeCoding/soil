package cruise_setting

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"time"
)

func New(
	sessionIdentifier string,
	mode constant.Cruise,
	pace time.Duration,
) *Setting {
	return &Setting{
		SessionIdentifier: sessionIdentifier,
		Mode:              mode,
		Pace:              pace,
	}
}
