package alfa

type Server struct {
	Name string
}

func New(name string) *Server {
	return &Server{Name: name}
}

func Named(name string) *Server {
	return &Server{Name: name}
}

func Wrap(s *Server) *Server {
	return s
}

func Keep(s *Server) *Server {
	return s
}
