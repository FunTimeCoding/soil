package event

import (
	"github.com/funtimecoding/soil/pkg/chromium/constant"
	"time"
)

func New(
	kind constant.EventKind,
	targetIdentifier string,
	locator string,
) *Event {
	return &Event{
		Kind:             kind,
		TargetIdentifier: targetIdentifier,
		Locator:          locator,
		At:               time.Now(),
	}
}
