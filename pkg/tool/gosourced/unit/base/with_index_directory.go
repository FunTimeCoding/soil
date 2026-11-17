package base

func WithIndexDirectory(directory string) Option {
	return func(s *Server) {
		s.indexDirectory = directory
	}
}
