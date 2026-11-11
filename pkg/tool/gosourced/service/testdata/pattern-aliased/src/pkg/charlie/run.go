package charlie

import (
	stepper "example/pkg/alfa"
	stepConstant "example/pkg/delta/constant"
)

func Run(n int) int {
	return stepper.Advance(n, stepConstant.Step)
}
