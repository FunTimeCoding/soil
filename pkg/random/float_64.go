package random

import "math/rand"

func (*Source) Float64() float64 {
	return rand.Float64()
}
