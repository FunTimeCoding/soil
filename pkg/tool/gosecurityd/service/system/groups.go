package system

import (
	"cmp"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/finding_group"
	"slices"
)

func (s *Service) Groups(m float64) ([]*finding_group.Group, error) {
	if e := minimum.Check(m); e != nil {
		return nil, e
	}

	details, e := s.Details()

	if e != nil {
		return nil, e
	}

	byVulnerability := map[string]*finding_group.Group{}

	for _, d := range details {
		for name, f := range worstEach(d.Findings) {
			if f.Score < m {
				continue
			}

			g, okay := byVulnerability[name]

			if !okay {
				g = finding_group.New(name, f.Score, f.Severity)
				byVulnerability[name] = g
			}

			g.Systems = append(g.Systems, d.Summary.Name)
			g.Count.Add(f.Kind())
			g.Unimportant = g.Unimportant && f.Urgency == constant.Unimportant
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
				cmp.Compare(b.Score, a.Score),
				cmp.Compare(len(b.Systems), len(a.Systems)),
				cmp.Compare(a.Vulnerability, b.Vulnerability),
			)
		},
	)

	return result, nil
}
