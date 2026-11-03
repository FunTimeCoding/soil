package gofix

import "github.com/funtimecoding/soil/pkg/source/types/import_decision"

func newAliasMode(
	suggest bool,
	applies func([]*import_decision.Decision) bool,
	ours []string,
) *aliasMode {
	return &aliasMode{suggest: suggest, applies: applies, ours: ours}
}
