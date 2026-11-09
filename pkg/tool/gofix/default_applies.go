package gofix

import (
	sourceConstant "github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/imports"
	"github.com/funtimecoding/soil/pkg/source/types/import_decision"
)

func defaultApplies(group []*import_decision.Decision) bool {
	marked := false
	renamed := false
	roots := true

	for _, d := range group {
		if d.Kind != sourceConstant.DecisionKeep &&
			d.Kind != sourceConstant.DecisionRename {
			return false
		}

		if imports.IsVersion(d.Entry.Declared) ||
			imports.IsNumberedFallback(d.Entry.Local, d.Entry.Declared) {
			marked = true
		}

		if d.Kind != sourceConstant.DecisionRename {
			continue
		}

		renamed = true

		if !imports.IsLayoutRoot(d.Entry.Path) {
			roots = false
		}
	}

	return marked || renamed && roots
}
