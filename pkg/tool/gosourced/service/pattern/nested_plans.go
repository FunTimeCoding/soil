package pattern

import "github.com/funtimecoding/soil/pkg/tool/gosourced/types/pattern_site"

func nestedPlans(plans []pattern_site.ApplyPlan) map[int]bool {
	result := map[int]bool{}

	for i, outer := range plans {
		for j, inner := range plans {
			if i == j {
				continue
			}

			if outer.Statement.Pos() <= inner.Statement.Pos() &&
				inner.Statement.End() <= outer.Statement.End() {
				result[i] = true
				result[j] = true
			}
		}
	}

	return result
}
