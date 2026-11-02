package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/types/change"
	"time"
)

func (s *Service) Publish(x context.Context) (string, []*change.Change, error) {
	pending, e := s.store.Unpublished()

	if e != nil {
		return "", nil, e
	}

	change, f := s.publisher.Changes(pending)

	if f != nil {
		return "", nil, f
	}

	if len(change) == 0 {
		return "", nil, nil
	}

	var commit string
	g := s.guard.Attempt(
		x,
		func(y context.Context) error {
			var h error
			commit, h = s.publisher.Commit(y, change, message(pending))

			return h
		},
	)

	if g != nil {
		return "", nil, g
	}

	var serial []string

	for _, r := range pending {
		serial = append(serial, r.Serial)
	}

	return commit, change, s.store.MarkPublished(serial, time.Now())
}
