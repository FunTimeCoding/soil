package example

import "io"

func ClosedInClosure() {
	c := newChained()
	defer func() { c.Close() }()
	c.Work()
}

func HandedClosure(register func(func())) {
	c := newChained()
	register(func() { c.Work() })
}

func ReturnedLocalClosure() func() {
	c := newChained()
	work := func() { c.Work() }

	return work
}

type Hook struct {
	run func()
}

func StoredClosure(h *Hook) {
	c := newChained()
	h.run = func() { c.Work() }
}

type Flusher struct {
	chained *Chained
}

func newFlusher() *Flusher {
	return &Flusher{chained: newChained()}
}

func (f *Flusher) Close() {
	f.chained.Close()
}

func (f *Flusher) Flush(any) {
	f.chained.Work()
	f.chained.Close()
}

func FlushedInDeferredClosure() {
	f := newFlusher()
	defer func() { f.Flush(recover()) }()
}

func PipeClosedInGoroutine(failure error) {
	_, writer := io.Pipe()
	go func() { _ = writer.CloseWithError(failure) }()
}
