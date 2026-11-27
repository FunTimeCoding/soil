package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/user"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gographd/constant"
	"strings"
)

func (s *Service) resolveUser(input string) (*user.User, error) {
	if strings.TrimSpace(input) == "" {
		return nil, validation.New("user is required")
	}

	if strings.Contains(input, "@") || constant.UserPattern.MatchString(input) {
		return peer.Guarded(
			context.Background(),
			s.guard,
			func(context.Context) (*user.User, error) {
				return s.client.User(input)
			},
		)
	}

	found, e := s.SearchUsers(input, constant.UserSearchLimit)

	if e != nil {
		return nil, e
	}

	var exact []*user.User

	for _, u := range found {
		if strings.EqualFold(u.Name, input) {
			exact = append(exact, u)
		}
	}

	switch {
	case len(exact) == 1:
		return exact[0], nil
	case len(exact) == 0 && len(found) == 1:
		return found[0], nil
	case len(found) == 0:
		return nil, not_found.New("user", input)
	default:
		return nil, ambiguous.Format(
			"user name matches %d users, use the email: %s",
			len(found),
			input,
		)
	}
}
