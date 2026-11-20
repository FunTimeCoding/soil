package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/status"
)

func Status(v *status.Status) *server.Status {
	result := &server.Status{
		Package:       v.Package,
		Release:       v.Release,
		Status:        v.Status,
		NeverAffected: v.NeverAffected,
		Urgency:       v.Urgency,
		Repositories:  v.Repositories,
	}

	if !v.NeverAffected {
		result.FixedVersion = optional(v.Fixed)
	}

	return result
}
