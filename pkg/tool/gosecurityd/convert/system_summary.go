package convert

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
)

func SystemSummary(v *system_result.Summary) *server.SystemSummary {
	result := &server.SystemSummary{
		Name:           v.Name,
		Release:        v.Release,
		Kernel:         optional(v.Kernel),
		Priority:       v.Priority,
		Coverage:       v.Coverage,
		Checked:        v.Checked,
		Clean:          v.Clean,
		Findings:       v.Findings,
		NoFix:          v.NoFix,
		PaidOnly:       v.PaidOnly,
		Upgradable:     v.Upgradable,
		Updated:        v.Updated,
		Stale:          v.Stale,
		TrackerStale:   v.TrackerStale,
		Reboot:         v.Reboot,
		RunningKernel:  optional(v.RunningKernel),
		NewestKernel:   optional(v.NewestKernel),
		KernelState:    v.KernelState,
		KernelReason:   optional(v.KernelReason),
		KernelFindings: v.KernelFindings,
		KernelWithFix:  v.KernelWithFix,
		Critical:       v.Critical,
		TopScore:       v.TopScore,
		ListsUpdated:   v.ListsUpdated,
	}

	if len(v.KernelOffers) > 0 {
		offers := make([]server.KernelOffer, len(v.KernelOffers))

		for i, o := range v.KernelOffers {
			offers[i] = server.KernelOffer{
				Version:  o.Version,
				Priority: o.Priority,
				Origins:  o.Origins,
			}
		}

		result.KernelOffers = &offers
	}

	if len(v.Labels) > 0 {
		result.Labels = &v.Labels
	}

	if !v.TrackerVerified.IsZero() {
		result.TrackerVerified = &v.TrackerVerified
	}

	return result
}
