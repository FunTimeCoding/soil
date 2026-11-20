package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func SystemFinding(v *system_result.Finding) *server.SystemFinding {
	return &server.SystemFinding{
		Vulnerability: v.Vulnerability,
		Package:       v.Package,
		Source:        v.Source,
		Installed:     v.Installed,
		Fixed:         optional(v.Fixed),
		Fix:           v.Fix,
		Urgency:       v.Urgency,
		Upgradable:    v.Upgradable,
	}
}
