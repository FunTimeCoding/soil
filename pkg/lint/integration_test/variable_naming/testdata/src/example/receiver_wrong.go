package example

type ReceiverTarget struct{}

func (h *ReceiverTarget) DoThing() {
	_ = h
}
