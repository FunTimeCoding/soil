package diff

import (
	goModConstant "github.com/funtimecoding/soil/pkg/go_mod/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/module_range"
)

func (s *Service) companions(
	self string,
	before map[string][]byte,
	after map[string][]byte,
	to string,
) []*module_range.Range {
	was := requirements(before[goModConstant.ModFile])
	now := requirements(after[goModConstant.ModFile])
	var result []*module_range.Range

	for _, m := range s.inventory.Modules {
		if m.Name == self {
			continue
		}

		p := modulePathOnDisk(m.Directory)
		from, inWas := was[p]
		next, inNow := now[p]

		if p == "" || !inWas || !inNow || from == "" {
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
