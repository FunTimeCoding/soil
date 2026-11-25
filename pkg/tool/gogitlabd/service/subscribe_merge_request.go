package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_detail"
	goclaudedConstant "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Service) SubscribeMergeRequest(
	x context.Context,
	instance string,
	project int64,
	request int64,
	callsign string,
	alias string,
	d goclaudedConstant.Delivery,
) (
	*subscription.Subscription,
	*merge_request_detail.Detail,
	bool,
	error,
) {
	if e := s.requireWatches(); e != nil {
		return nil, nil, false, e
	}

	k := watched_request.New(instance, project, request)
	existing, found, e := s.store.Find(callsign, k)

	if e != nil {
		return nil, nil, false, e
	}

	if found {
		return existing, nil, true, nil
	}

	current, detail, f := s.observe(x, k)

	if f != nil {
		return nil, nil, false, f
	}

	if current.Finished() {
		return nil, nil, false, validation.New(
			"%s is %s, nothing to watch",
			reference(detail),
			current.State,
		)
	}

	_, known, g := s.store.Snapshot(k)

	if g != nil {
		return nil, nil, false, g
	}

	if !known {
		if h := s.baseline(x, k, current); h != nil {
			return nil, nil, false, h
		}
	}

	result := subscription.New(callsign, k, reference(detail), alias, d)

	if h := s.store.Create(result); h != nil {
		return nil, nil, false, h
	}

	return result, detail, false, nil
}
