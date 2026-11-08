package parameter

type Busy struct {
	value int
}

func (y *Busy) Scale(s int) int {
	return y.value * s
}
