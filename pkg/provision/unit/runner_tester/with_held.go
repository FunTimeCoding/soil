package runner_tester

func WithHeld() Option {
	return func(s *setup) {
		gate := make(chan struct{})
		s.gate = gate
		s.hold = func() { <-gate }
	}
}
