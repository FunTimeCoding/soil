package import_decision

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/types/import_entry"
)

func New(
	entry *import_entry.Entry,
	rule constant.DecisionRule,
) *Decision {
	return &Decision{
		Entry: entry,
		Want:  entry.Local,
		Kind:  constant.DecisionKeep,
		Rule:  rule,
	}
}
