package service

import (
	"github.com/funtimecoding/soil/pkg/chat/constant"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/post"
	"github.com/funtimecoding/soil/pkg/chat/mattermost/thread"
)

func (s *Service) unresolved(p *post.Post) (*thread.Thread, error) {
	if p.Raw.RootId != "" {
		return nil, nil
	}

	done, e := s.mattermost.PostHasReaction(
		p.Raw,
		constant.MattermostDoneReaction,
	)

	if e != nil || done {
		return nil, e
	}

	replies, f := s.mattermost.Thread(p.Raw)

	if f != nil {
		return nil, f
	}

	return thread.New(p, replies), nil
}
