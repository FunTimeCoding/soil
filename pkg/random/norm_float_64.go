package random

import "math/rand"

func (*Source) NormFloat64() float64 {
	return rand.NormFloat64()
}
