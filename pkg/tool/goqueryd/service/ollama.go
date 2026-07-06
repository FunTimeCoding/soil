package service

import "github.com/funtimecoding/soil/pkg/generative/ollama"

func (s *Service) Ollama() *ollama.Client {
	return s.ollama
}
