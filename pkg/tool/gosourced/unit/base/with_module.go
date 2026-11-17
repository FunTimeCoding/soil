package base

func WithModule(
	name string,
	directory string,
) Option {
	return func(s *Server) {
		s.inventory.Add(name, directory)
	}
}
