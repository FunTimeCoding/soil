package bravo

import (
	"example/pkg/alfa"
	deltaConstant "example/pkg/delta/constant"
)

func Run(n int) int {
	return alfa.Advance(n, deltaConstant.Step)
}
