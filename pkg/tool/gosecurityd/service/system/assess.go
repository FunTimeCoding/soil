package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (s *Service) assess(v *system.System) (*system_result.Detail, error) {
	m, e := s.findings(v)

	if e != nil {
		return nil, e
	}

	findings := m.Findings
	now := s.clock()
	summary := system_result.NewSummary(v)
	summary.Stale = now.Sub(v.UpdatedAt) > s.staleAfter
	summary.RunningKernel = m.RunningKernel
	summary.NewestKernel = m.NewestKernel
	summary.KernelState = kernelState(m)

	if slices.ContainsFunc(
		findings,
		(*system_result.Finding).UnreachableKernel,
	) {
		summary.KernelReason = kernelReason(v, m.RunningKernel, now)
	}

	if source := s.debian.Covering(v.Release); source != nil {
		summary.Coverage = constant.CoverageTracked

		if source.Supported() {
			summary.Coverage = constant.CoverageSupported
		}

		if source.Tracker() != nil {
			summary.Checked = true
			summary.TrackerVerified = source.Verified()
			summary.TrackerStale = now.Sub(
				summary.TrackerVerified,
			) > s.trackerStaleAfter
		}
	}

	critical := map[string]bool{}

	for _, f := range findings {
		summary.TopScore = max(summary.TopScore, f.Score)

		if f.Score >= constant.CriticalScore {
			critical[f.Vulnerability] = true
		}

		if f.Kernel {
			summary.KernelFindings++

			if f.Fix != constant.FindingNone {
				summary.KernelWithFix++
			}

			if f.Reboot {
				summary.Reboot++
			}

			continue
		}

		summary.Findings++

		switch f.Fix {
		case constant.FindingNone:
			summary.NoFix++
		case constant.FindingPaid:
			summary.PaidOnly++
		}

		if f.Upgradable {
			summary.Upgradable++
		}
	}

	summary.Critical = len(critical)
	summary.Clean = summary.Coverage == constant.CoverageSupported &&
		summary.Checked &&
		!summary.Stale &&
		!summary.TrackerStale &&
		len(findings) == 0

	return system_result.NewDetail(summary, findings), nil
}
