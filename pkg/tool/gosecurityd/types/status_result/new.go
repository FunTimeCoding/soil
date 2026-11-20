package status_result

import (
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
	"time"
)

func New(
	identifier string,
	statuses []*status.Status,
	modified time.Time,
) *Result {
	return &Result{
		Identifier: identifier,
		Statuses:   statuses,
		Modified:   modified,
	}
}
