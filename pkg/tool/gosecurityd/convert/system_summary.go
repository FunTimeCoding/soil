package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func SystemSummary(v *system_result.Summary) *server.SystemSummary {
	result := &server.SystemSummary{
		Name:         v.Name,
		Release:      v.Release,
		Kernel:       optional(v.Kernel),
		Priority:     v.Priority,
		Coverage:     v.Coverage,
		Checked:      v.Checked,
		Clean:        v.Clean,
		Findings:     v.Findings,
		NoFix:        v.NoFix,
		PaidOnly:     v.PaidOnly,
		Upgradable:   v.Upgradable,
		Updated:      v.Updated,
		Stale:        v.Stale,
		TrackerStale: v.TrackerStale,
	}

	if len(v.Labels) > 0 {
		result.Labels = &v.Labels
	}

	if !v.TrackerVerified.IsZero() {
		result.TrackerVerified = &v.TrackerVerified
	}

	return result
}
