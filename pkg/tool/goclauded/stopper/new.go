package stopper

func New() *Stopper {
	return &Stopper{done: make(chan struct{})}
}
