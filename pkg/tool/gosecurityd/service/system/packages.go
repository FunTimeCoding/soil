package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/package_group"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/package_result"
	"github.com/funtimecoding/soil/pkg/vulnerability/debian/version"
)

func (s *Service) Packages(name string) (*package_result.Result, error) {
	d, e := s.Detail(name)

	if e != nil {
		return nil, e
	}

	result := package_result.New(d.Summary)
	groups := map[*[]*package_group.Group]map[string]*package_group.Group{}

	for _, f := range d.Findings {
		list := &result.Open

		switch {
		case f.Upgradable:
			list = &result.Upgrade
		case f.Fix == constant.FindingDebian:
			list = &result.Pending
		case f.Fix == constant.FindingPaid:
			list = &result.Paid
		}

		if groups[list] == nil {
			groups[list] = map[string]*package_group.Group{}
		}

		g, okay := groups[list][f.Source]

		if !okay {
			g = package_group.New(f.Source, f.Installed, f.Candidate)
			groups[list][f.Source] = g
			*list = append(*list, g)
		}

		g.Vulnerabilities = append(g.Vulnerabilities, f.Vulnerability)

		if f.Fixed != "" &&
			(g.Fixed == "" || version.Compare(f.Fixed, g.Fixed) > 0) {
			g.Fixed = f.Fixed
		}
	}

	return result, nil
}
