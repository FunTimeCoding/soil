package status_result

import (
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
	"time"
)

type Result struct {
	Identifier string
	Statuses   []*status.Status
	Modified   time.Time
}
