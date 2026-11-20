package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/status_result"
)

func DebianStatus(v *status_result.Result) *server.DebianStatus {
	result := &server.DebianStatus{
		Identifier: v.Identifier,
		Statuses:   Statuses(v.Statuses),
	}

	if !v.Modified.IsZero() {
		result.TrackerModified = &v.Modified
	}

	return result
}
