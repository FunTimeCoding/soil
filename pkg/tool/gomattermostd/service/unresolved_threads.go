package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
	"github.com/funtimecoding/soil/pkg/peer"
	"time"
)

func (s *Service) UnresolvedThreads(
	name string,
	since time.Time,
) ([]*thread.Thread, error) {
	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]*thread.Thread, error) {
			h, e := s.mattermost.TeamChannel(name)

			if e != nil {
				return nil, e
			}

			posts, f := s.mattermost.Posts(h, since)

			if f != nil {
				return nil, f
			}

			var result []*thread.Thread

			for _, p := range posts {
				t, g := s.unresolved(p)

				if g != nil {
					return nil, g
				}

				if t != nil {
					result = append(result, t)
				}
			}

			return result, nil
		},
	)
}
