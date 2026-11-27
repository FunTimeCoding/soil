package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/teams/destination"
	"github.com/funtimecoding/soil/pkg/chat/teams/member"
	"github.com/funtimecoding/soil/pkg/errors/ambiguous"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/peer"
	"strings"
)

func (s *Service) resolveMembers(
	d *destination.Destination,
	names []string,
) ([]*member.Member, error) {
	if len(names) == 0 {
		return nil, nil
	}

	members, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*member.Member, error) {
			return s.client.Members(d)
		},
	)

	if e != nil {
		return nil, e
	}

	var result []*member.Member

	for _, n := range names {
		var matches []*member.Member

		for _, m := range members {
			if strings.EqualFold(m.Name, n) {
				matches = append(matches, m)
			}
		}

		switch len(matches) {
		case 0:
			return nil, not_found.New("member", n)
		case 1:
			result = append(result, matches[0])
		default:
			return nil, ambiguous.Format(
				"member name matches %d members: %s",
				len(matches),
				n,
			)
		}
	}

	return result, nil
}
