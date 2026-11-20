package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/installed"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/store/system"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"slices"
)

func (s *Service) assess(
	v *system.System,
	packages []installed.Installed,
) *system_result.Detail {
	now := s.clock()
	summary := system_result.NewSummary(v)
	summary.Stale = now.Sub(v.UpdatedAt) > s.staleAfter
	source := s.debian.Covering(v.Release)
	var findings []*system_result.Finding

	if source != nil {
		summary.Coverage = constant.CoverageTracked

		if source.Supported() {
			summary.Coverage = constant.CoverageSupported
		}

		if t := source.Tracker(); t != nil {
			summary.Checked = true
			summary.TrackerVerified = source.Verified()
			summary.TrackerStale = now.Sub(
				summary.TrackerVerified,
			) > s.trackerStaleAfter
			seen := map[[2]string]bool{}

			for _, p := range packages {
				for _, f := range match(t, v.Release, &p, source.Supported()) {
					k := [2]string{f.Vulnerability, f.Source}

					if !seen[k] {
						seen[k] = true
						findings = append(findings, f)
					}
				}
			}
		}
	}

	for _, f := range findings {
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

	summary.Findings = len(findings)
	summary.Clean = summary.Coverage == constant.CoverageSupported &&
		summary.Checked &&
		!summary.Stale &&
		!summary.TrackerStale &&
		len(findings) == 0
	slices.SortFunc(
		findings,
		func(a *system_result.Finding, b *system_result.Finding) int {
			return cmp.Or(
				cmp.Compare(a.Source, b.Source),
				cmp.Compare(a.Vulnerability, b.Vulnerability),
			)
		},
	)

	return system_result.NewDetail(summary, findings)
}
