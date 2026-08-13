package notifier

func New() *Notifier {
	return &Notifier{subscribers: make(map[chan struct{}]struct{})}
}
