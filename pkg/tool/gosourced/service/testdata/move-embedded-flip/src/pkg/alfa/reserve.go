package alfa

type reserve struct {
	Amount int
}

type Holder struct {
	reserve
}

func New() Holder {
	return Holder{reserve: reserve{Amount: 1}}
}

func (h Holder) Total() int {
	return h.reserve.Amount
}

func Mover() int {
	return reserve{Amount: 3}.Amount
}
