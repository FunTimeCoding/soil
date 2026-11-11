package bravo

import "example/pkg/alfa"

type Holder struct {
	alfa.Reserve
	Name string
}

func New() *Holder {
	return &Holder{Reserve: alfa.Reserve{Amount: 1}, Name: "bravo"}
}

func (h *Holder) Total() int {
	return h.Reserve.Amount
}
