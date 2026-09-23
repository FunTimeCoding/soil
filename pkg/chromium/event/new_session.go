package event

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"time"
)

func NewSession(
	kind constant.EventKind,
	sessionIdentifier string,
) *Event {
	return &Event{
		Kind:              kind,
		SessionIdentifier: sessionIdentifier,
		At:                time.Now(),
	}
}
