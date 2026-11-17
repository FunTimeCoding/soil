package service_tester

func WithDirectory(directory string) Option {
	return func(s *setup) {
		s.directory = directory
	}
}
