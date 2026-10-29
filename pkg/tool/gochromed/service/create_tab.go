package service

import "github.com/funtimecoding/soil/pkg/errors/validation"

func (s *Service) CreateTab(l string) (string, error) {
	if l == "" {
		return "", validation.New("url is required")
	}

	return s.source.CreateTab(l)
}
