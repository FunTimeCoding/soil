package event

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"time"
)

func New(
	kind constant.ChromiumEventKind,
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
