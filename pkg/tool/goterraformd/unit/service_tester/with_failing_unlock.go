package service_tester

func WithFailingUnlock() Option {
	return func(o *Tester) {
		o.unlockExit = 1
	}
}
