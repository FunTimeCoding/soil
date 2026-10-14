package fact

import "github.com/funtimecoding/soil/pkg/lint/constant"

func NewCondition(
	callee string,
	argument int,
) *Condition {
	if callee == "" {
		return &Condition{}
	}

	return &Condition{
		Kind:     constant.DisposesCondition,
		Callee:   callee,
		Argument: argument,
	}
}
