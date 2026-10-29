package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"
	"github.com/mattermost/mattermost/server/public/model"
)

func (s *Service) Subscribe(
	p *model.Post,
	callsign string,
	alias string,
) (*subscription.Subscription, bool, error) {
	if e := s.requireSubscriptions(); e != nil {
		return nil, false, e
	}

	root := p.Id

	if p.RootId != "" {
		root = p.RootId
	}

	existing, e := s.store.ByRoot(root)

	if e != nil {
		return nil, false, e
	}

	for _, v := range existing {
		if v.Callsign == callsign {
			return &v, true, nil
		}
	}

	result := subscription.New(callsign, root, p.ChannelId, alias)

	if f := s.store.Create(result); f != nil {
		return nil, false, f
	}

	s.watcher.Index(root)

	return result, false, nil
}
