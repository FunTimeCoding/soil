package event

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"time"
)

type Event struct {
	Kind   constant.EventKind
	Author string
	Text   string
	At     time.Time
}
