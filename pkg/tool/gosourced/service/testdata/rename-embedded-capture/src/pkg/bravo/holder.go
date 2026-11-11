package bravo

import "example/pkg/alfa"

type Extra struct {
	Budget int
}

type Holder struct {
	alfa.Reserve
	Extra
}

func Read(h *Holder) int {
	return h.Budget
}
