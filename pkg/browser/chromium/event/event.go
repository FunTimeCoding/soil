package event

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"time"
)

type Event struct {
	Kind              constant.ChromiumEventKind
	TargetIdentifier  string
	SessionIdentifier string
	Locator           string
	At                time.Time
}
