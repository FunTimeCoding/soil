package service

import (
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
	"time"
)

func (s *Service) UnresolvedThreads(
	name string,
	since time.Time,
) ([]*thread.Thread, error) {
	return guarded(
		s.guard,
		func() ([]*thread.Thread, error) {
			h, e := s.mattermost.TeamChannel(name)

			if e != nil {
				return nil, e
			}

			posts, f := s.mattermost.PostsSince(h, since)

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
