package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
	"time"
)

func (s *Service) PollWatches(x context.Context) []error {
	if s.store == nil {
		return nil
	}

	all, e := s.store.All()

	if e != nil {
		return []error{e}
	}

	now := time.Now()
	due := map[string]bool{}

	for _, root := range s.pending.Roots(now) {
		due[root] = true
	}

	var order []*watched_request.Key
	groups := map[string][]subscription.Subscription{}

	for _, v := range all {
		k := v.Key()

		if _, okay := groups[k.String()]; !okay {
			order = append(order, k)
		}

		groups[k.String()] = append(groups[k.String()], v)
	}

	var result []error

	for _, k := range order {
		result = append(
			result,
			s.pollRequest(x, k, groups[k.String()], due[k.String()], now)...,
		)
	}

	s.writers.Prune(now.Add(-constant.WriterWindow))

	return result
}
