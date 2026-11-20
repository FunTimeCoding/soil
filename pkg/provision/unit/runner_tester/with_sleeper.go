package runner_tester

func WithSleeper() Option {
	return func(s *setup) {
		s.sleeper = true
	}
}
