package echo

import (
	"example/pkg/alfa"
	"example/pkg/delta"
)

func Use() int {
	delta := delta.Existing()

	return alfa.Shifting() + delta
}
