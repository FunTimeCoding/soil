package alfa

type Server struct {
	Port int
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() {}

func restart() {
	NewServer().Start()
}
