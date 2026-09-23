package cruise

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"time"
)

func New(
	session string,
	mode constant.Cruise,
	pace time.Duration,
) *Cruise {
	return &Cruise{Session: session, Mode: mode, Pace: pace}
}
