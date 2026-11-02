package example

type Chained struct{}

func newChained() *Chained {
	return &Chained{}
}

func (c *Chained) With() *Chained {
	return c
}

func (c *Chained) Work() {}

func (c *Chained) Close() {}

func ClosedChain() {
	newChained().Close()
}

func DeferredChain() {
	defer newChained().Close()
}

func BuiltChain() {
	c := newChained().With()
	defer c.Close()
	c.Work()
}

func ReturnedChain() *Chained {
	return newChained().With()
}

type Lender struct {
	inner *Chained
}

func newLender() *Lender {
	return &Lender{inner: newChained()}
}

func (l *Lender) Inner() *Chained {
	return l.inner
}

func (l *Lender) Close() {
	l.inner.Close()
}

func Lent() *Chained {
	return newLender().Inner()
}

type Holder struct {
	chained *Chained
}

func (h *Holder) View() *Chained {
	return &Chained{}
}

func DerivedFromHeld(h *Holder) {
	h.View().Work()
}
