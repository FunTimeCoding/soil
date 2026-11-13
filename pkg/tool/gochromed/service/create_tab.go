package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) CreateTab(l string) (string, error) {
	if l == "" {
		return "", validation.New("url is required")
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (string, error) {
			return s.source.CreateTab(l)
		},
	)
}
