package service

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"

func (s *Service) MergeRequestSubscriptions(
	callsign string,
) ([]subscription.Subscription, error) {
	if e := s.requireWatches(); e != nil {
		return nil, e
	}

	return s.store.ByCallsign(callsign)
}
