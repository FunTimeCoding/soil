package event

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"time"
)

type Event struct {
	Kind              constant.EventKind
	TargetIdentifier  string
	SessionIdentifier string
	Locator           string
	At                time.Time
}
