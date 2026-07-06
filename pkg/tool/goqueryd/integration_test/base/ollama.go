package base

import "github.com/funtimecoding/soil/pkg/generative/ollama"

func (s *Server) Ollama() *ollama.Client {
	return s.ollama
}
