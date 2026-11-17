package base

func WithSeedDirectory(directory string) Option {
	return func(o *Server) {
		o.seedDirectory = directory
	}
}
