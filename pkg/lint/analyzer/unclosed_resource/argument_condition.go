package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func argumentCondition(
	f *types.Func,
	index int,
) *fact.Condition {
	return fact.NewCondition(functionKey(f), index)
}
