package gofix

import (
	sourceConstant "github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
)

func appliesWhole(group []*import_decision.Decision) bool {
	for _, d := range group {
		if d.Kind != sourceConstant.DecisionKeep &&
			d.Kind != sourceConstant.DecisionRename {
			return false
		}
	}

	return true
}
