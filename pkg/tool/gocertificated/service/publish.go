package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/publish"
	"time"
)

func (s *Service) Publish() (string, []*publish.Change, error) {
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

	commit, g := s.publisher.Commit(change, message(pending))

	if g != nil {
		return "", nil, g
	}

	var serial []string

	for _, r := range pending {
		serial = append(serial, r.Serial)
	}

	return commit, change, s.store.MarkPublished(serial, time.Now())
}
