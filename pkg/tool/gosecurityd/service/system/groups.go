package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/finding_group"
	"slices"
)

func (s *Service) Groups() ([]*finding_group.Group, error) {
	details, e := s.details()

	if e != nil {
		return nil, e
	}

	byVulnerability := map[string]*finding_group.Group{}

	for _, d := range details {
		counted := map[string]bool{}

		for _, f := range d.Findings {
			g, okay := byVulnerability[f.Vulnerability]

			if !okay {
				g = finding_group.New(f.Vulnerability)
				byVulnerability[f.Vulnerability] = g
			}

			if counted[f.Vulnerability] {
				continue
			}

			counted[f.Vulnerability] = true
			g.Systems = append(g.Systems, d.Summary.Name)

			switch f.Fix {
			case constant.FindingNone:
				g.NoFix++
			case constant.FindingPaid:
				g.PaidOnly++
			}

			if f.Upgradable {
				g.Upgradable++
			}
		}
	}

	result := make([]*finding_group.Group, 0, len(byVulnerability))

	for _, g := range byVulnerability {
		result = append(result, g)
	}

	slices.SortFunc(
		result,
		func(a *finding_group.Group, b *finding_group.Group) int {
			return cmp.Or(
				cmp.Compare(len(b.Systems), len(a.Systems)),
				cmp.Compare(a.Vulnerability, b.Vulnerability),
			)
		},
	)

	return result, nil
}
