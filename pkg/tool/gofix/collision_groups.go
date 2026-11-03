package gofix

import (
	sourceConstant "github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
)

func collisionGroups(
	decisions []*import_decision.Decision,
) [][]*import_decision.Decision {
	index := make(map[string]int)
	var result [][]*import_decision.Decision

	for _, d := range decisions {
		if d.Rule != sourceConstant.RuleCollision {
			continue
		}

		i, okay := index[d.Entry.Declared]

		if !okay {
			i = len(result)
			index[d.Entry.Declared] = i
			result = append(result, nil)
		}

		result[i] = append(result[i], d)
	}

	return result
}
