package fact

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strconv"
)

func NewRelayCondition(
	source string,
	argument int,
	sink string,
	sinkArgument int,
) *Condition {
	return &Condition{
		Kind:     constant.RelaysCondition,
		Callee:   source,
		Argument: argument,
		Holder:   sink,
		Field:    strconv.Itoa(sinkArgument),
	}
}
