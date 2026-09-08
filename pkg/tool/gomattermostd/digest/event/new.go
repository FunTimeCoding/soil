package event

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"time"
)

func New(
	kind constant.EventKind,
	author string,
	text string,
	at time.Time,
) *Event {
	return &Event{Kind: kind, Author: author, Text: text, At: at}
}
