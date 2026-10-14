package face

type Drainer interface {
	Finished() <-chan struct{}
	Interrupt()
}
