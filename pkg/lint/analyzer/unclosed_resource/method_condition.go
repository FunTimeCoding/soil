package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/types"
)

func methodCondition(f *types.Func) *fact.Condition {
	if isCloseMethod(f) {
		return fact.NewCondition("", 0)
	}

	return fact.NewCondition(functionKey(f), constant.ReceiverArgument)
}
