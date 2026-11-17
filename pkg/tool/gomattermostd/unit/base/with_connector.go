package base

func WithConnector() Option {
	return func(s *setup) {
		s.connector = true
	}
}
