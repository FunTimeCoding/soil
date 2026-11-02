package fact

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewFieldDrawCondition(
	callee string,
	argument int,
	field string,
) *Condition {
	return &Condition{
		Kind:     constant.DrawsCondition,
		Callee:   callee,
		Argument: argument,
		Field:    field,
	}
}
