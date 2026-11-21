package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/minimum"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/package_group"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/package_result"
)

func (s *Service) Packages(
	name string,
	m float64,
) (*package_result.Result, error) {
	if e := minimum.Check(m); e != nil {
		return nil, e
	}

	d, e := s.Detail(name)

	if e != nil {
		return nil, e
	}

	result := package_result.New(d.Summary)
	lists := map[string]*[]*package_group.Group{
		constant.KindReboot:        &result.Reboot,
		constant.KindKernelUpgrade: &result.Upgrade,
		constant.KindUpgrade:       &result.Upgrade,
		constant.KindPending:       &result.Pending,
		constant.KindPaid:          &result.Paid,
		constant.KindOpen:          &result.Open,
	}
	groups := map[[3]string]*package_group.Group{}

	for _, f := range d.Findings {
		if f.Score < m {
			continue
		}

		k := f.Kind()
		list := lists[k]
		key := [3]string{k, f.Source, ""}

		if f.Kernel {
			list = &result.Kernel
			key[2] = constant.KernelSource
		}

		g, okay := groups[key]

		if !okay {
			g = package_group.New(f.Source, k, f.Installed, f.Candidate)
			groups[key] = g
			*list = append(*list, g)
		}

		g.Vulnerabilities = append(g.Vulnerabilities, f.Vulnerability)
		g.Fixed = highest(g.Fixed, f.Fixed)
		g.TopScore = max(g.TopScore, f.Score)
	}

	return result, nil
}
