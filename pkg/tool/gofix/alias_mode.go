package gofix

import "github.com/funtimecoding/soil/pkg/source/types/import_decision"

type aliasMode struct {
	suggest bool
	applies func([]*import_decision.Decision) bool
	ours    []string
}
