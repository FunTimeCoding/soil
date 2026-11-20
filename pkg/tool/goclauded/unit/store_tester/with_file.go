package store_tester

func WithFile(path string) Option {
	return func(s *setup) {
		s.path = path
	}
}
