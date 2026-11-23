package system

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/upgrade_group"
	"slices"
	"strings"
)

func (s *Service) Upgrade(
	release string,
	source string,
) (*upgrade_group.Group, error) {
	r, e := s.Upgrades(
		system_selection.New(),
		release,
		kernel(strings.TrimSuffix(source, constant.BackportsLabel)),
	)

	if e != nil {
		return nil, e
	}

	for _, g := range r.Groups {
		if g.Release != release || g.Source != source {
			continue
		}

		for _, host := range g.Hosts {
			v, f := s.store.SystemByName(host.Name)

			if f != nil {
				return nil, f
			}

			packages, h := s.store.Packages(v.Identifier)

			if h != nil {
				return nil, h
			}

			for _, p := range packages {
				if p.Candidate == "" {
					continue
				}

				if g.Kernel && strings.HasPrefix(p.Name, constant.KernelImagePrefix) ||
					!g.Kernel && p.Source == source {
					host.Binaries = append(host.Binaries, p.Name)
				}
			}

			slices.Sort(host.Binaries)
		}

		return g, nil
	}

	return nil, not_found.New(constant.UpgradeKind, join.Space(release, source))
}
