package runner_tester

func withHold(f func()) Option {
	return func(s *setup) {
		s.hold = f
	}
}
