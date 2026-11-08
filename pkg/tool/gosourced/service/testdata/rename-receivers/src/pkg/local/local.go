package local

type Clean struct {
	value int
}

func (y *Clean) Value() int {
	return y.value
}

type Busy struct {
	value int
}

func (y *Busy) Value() int {
	s := 2

	return y.value * s
}
