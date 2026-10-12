package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/operation"
)

func operations(available []int) []*operation.Operation {
	var result []*operation.Operation

	for _, name := range constant.Operations {
		forces := []bool{false}

		if name != constant.OperationStart {
			forces = append(forces, true)
		}

		for _, force := range forces {
			state, e := requestedState(name, force)
			errors.PanicOnError(e)
			var refusal string

			if _, f := Decide(name, force, available); f != nil {
				refusal = f.Error()
			}

			result = append(
				result,
				operation.New(name, force, band.PowerStateName(state), refusal),
			)
		}
	}

	return result
}
