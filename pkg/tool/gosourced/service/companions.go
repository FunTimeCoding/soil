package service

import (
	modConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/module_range"
)

func (s *Service) companions(
	self string,
	before map[string][]byte,
	after map[string][]byte,
	to string,
) []*module_range.Range {
	was := requirements(before[modConstant.ModFile])
	now := requirements(after[modConstant.ModFile])
	var result []*module_range.Range

	for _, m := range s.inventory.Modules {
		if m.Name == self {
			continue
		}

		p := modulePathOnDisk(m.Directory)
		from, inBefore := was[p]
		next, inAfter := now[p]

		if p == "" || !inBefore || !inAfter || from == "" {
			continue
		}

		if next == "" && to != "" {
			continue
		}

		if next != "" {
			if from == next {
				continue
			}

			next = versionRevision(next)
		}

		result = append(
			result,
			module_range.New(m.Name, p, versionRevision(from), next),
		)
	}

	return result
}
