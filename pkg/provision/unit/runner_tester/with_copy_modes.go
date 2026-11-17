package runner_tester

func WithCopyModes() Option {
	return func(s *setup) {
		s.copyModes = true
	}
}
