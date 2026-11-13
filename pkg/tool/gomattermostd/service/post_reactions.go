package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) PostReactions(identifier string) ([]string, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]string, error) {
			p, e := s.mattermost.FindPost(identifier)

			if e != nil {
				return nil, e
			}

			result := []string{}

			if !p.HasReactions {
				return result, nil
			}

			reactions, f := s.mattermost.Reactions(p)

			if f != nil {
				return nil, f
			}

			for _, v := range reactions {
				result = append(result, v.EmojiName)
			}

			return result, nil
		},
	)
}
