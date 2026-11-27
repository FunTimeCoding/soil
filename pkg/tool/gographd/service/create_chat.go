package service

import (
	"context"
	chatConstant "github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/teams/chat"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"slices"
)

func (s *Service) CreateChat(
	members []string,
	topic string,
) (*chat.Chat, error) {
	if len(members) == 0 {
		return nil, validation.New("chat members are required")
	}

	me, e := s.me()

	if e != nil {
		return nil, e
	}

	users := []string{me.Identifier}

	for _, m := range members {
		u, f := s.resolveUser(m)

		if f != nil {
			return nil, f
		}

		if !slices.Contains(users, u.Identifier) {
			users = append(users, u.Identifier)
		}
	}

	kind := chatConstant.TeamsGroupChat

	switch len(users) {
	case 1:
		return nil, validation.New("a chat needs someone besides yourself")
	case 2:
		if topic != "" {
			return nil, validation.New("a one-on-one chat has no topic")
		}

		kind = chatConstant.TeamsDirectChat
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*chat.Chat, error) {
			return s.client.CreateChat(kind, users, topic)
		},
	)
}
