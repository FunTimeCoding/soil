package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_host"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_result"
	"maps"
	"slices"
)

func (s *Service) Upgrades(
	selection *system_selection.Selection,
	release string,
	kernelOnly bool,
) (*upgrade_result.Result, error) {
	if release != "" &&
		!slices.Contains(constant.SupportedReleases, release) &&
		!slices.Contains(constant.ExtendedReleases, release) {
		return nil, validation.New("release unknown: %s", release)
	}

	m := selection.Minimum

	if e := minimum.Check(m); e != nil {
		return nil, e
	}

	details, f := s.Selected(selection.WithMinimum(0))

	if f != nil {
		return nil, f
	}

	result := upgrade_result.New()
	groups := map[[2]string]*upgrade_group.Group{}
	vulnerabilities := map[[2]string]map[string]bool{}
	releases := map[string]bool{}
	hosts := map[string]bool{}
	fixed := map[string]bool{}

	for _, d := range details {
		for _, f := range d.Findings {
			if !f.Upgradable || f.Reboot {
				continue
			}

			releases[d.Summary.Release] = true

			if f.Kernel != kernelOnly ||
				(release != "" && d.Summary.Release != release) ||
				f.Score < m {
				continue
			}

			label := f.Label()
			k := [2]string{d.Summary.Release, label}
			g, okay := groups[k]

			if !okay {
				g = upgrade_group.New(d.Summary.Release, label, kernelOnly)
				groups[k] = g
				vulnerabilities[k] = map[string]bool{}
				result.Groups = append(result.Groups, g)
			}

			if n := len(g.Hosts); n == 0 || g.Hosts[n-1].Name != d.Summary.Name {
				g.Hosts = append(g.Hosts, upgrade_host.New(d.Summary, f))
				hosts[d.Summary.Name] = true
				g.InstalledLowest = lowest(g.InstalledLowest, f.Installed)
				g.InstalledHighest = highest(g.InstalledHighest, f.Installed)
				g.Candidate = highest(g.Candidate, f.Candidate)
			}

			if !vulnerabilities[k][f.Vulnerability] {
				vulnerabilities[k][f.Vulnerability] = true
				g.Vulnerabilities = append(g.Vulnerabilities, f.Vulnerability)
			}

			g.Fixed = highest(g.Fixed, f.Fixed)
			g.TopScore = max(g.TopScore, f.Score)
			fixed[f.Vulnerability] = true
		}
	}

	for _, g := range result.Groups {
		slices.Sort(g.Vulnerabilities)
		slices.SortFunc(
			g.Hosts,
			func(a *upgrade_host.Host, b *upgrade_host.Host) int {
				return cmp.Or(
					cmp.Compare(a.Priority, b.Priority),
					cmp.Compare(a.Name, b.Name),
				)
			},
		)
	}

	slices.SortFunc(
		result.Groups,
		func(a *upgrade_group.Group, b *upgrade_group.Group) int {
			return cmp.Or(
				cmp.Compare(
					len(b.Hosts)*len(b.Vulnerabilities),
					len(a.Hosts)*len(a.Vulnerabilities),
				),
				cmp.Compare(len(b.Hosts), len(a.Hosts)),
				cmp.Compare(a.Source, b.Source),
				cmp.Compare(a.Release, b.Release),
			)
		},
	)
	result.Hosts = len(hosts)
	result.Vulnerabilities = len(fixed)
	result.Releases = slices.Sorted(maps.Keys(releases))

	return result, nil
}
