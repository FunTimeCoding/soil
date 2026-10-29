package service

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Service) Subscriptions(
	callsign string,
) ([]subscription.Subscription, error) {
	if e := s.requireSubscriptions(); e != nil {
		return nil, e
	}

	return s.store.ByCallsign(callsign)
}
