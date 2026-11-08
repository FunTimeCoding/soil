package outer

var s = &Busy{value: 3}

type Busy struct {
	value int
}

func (y *Busy) Plus() int {
	return y.value + s.value
}
