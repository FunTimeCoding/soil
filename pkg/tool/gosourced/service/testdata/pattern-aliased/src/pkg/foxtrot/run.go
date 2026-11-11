package foxtrot

import (
	"example/pkg/alfa"
	deltaConstant "example/pkg/echo/constant"
)

func Run(n int) int {
	return alfa.Advance(n, deltaConstant.Step)
}
