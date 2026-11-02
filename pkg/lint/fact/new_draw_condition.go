package fact

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewDrawCondition(
	callee string,
	argument int,
) *Condition {
	return &Condition{
		Kind:     constant.DrawsCondition,
		Callee:   callee,
		Argument: argument,
		Field:    constant.AnyField,
	}
}
