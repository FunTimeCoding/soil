package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"strings"
)

func (s *Service) SearchUsers(
	query string,
	limit int,
) ([]*user.User, error) {
	if strings.TrimSpace(query) == "" {
		return nil, validation.New("search query is required")
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*user.User, error) {
			return s.client.Users(query, limit)
		},
	)
}
