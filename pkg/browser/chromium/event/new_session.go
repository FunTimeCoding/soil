package event

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"time"
)

func NewSession(
	kind constant.ChromiumEventKind,
	sessionIdentifier string,
) *Event {
	return &Event{
		Kind:              kind,
		SessionIdentifier: sessionIdentifier,
		At:                time.Now(),
	}
}
