package fact

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewWrapCondition(
	callee string,
	argument int,
) *Condition {
	return &Condition{
		Kind:     constant.WrapsCondition,
		Callee:   callee,
		Argument: argument,
	}
}
