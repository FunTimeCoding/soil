package service

import "github.com/funtimecoding/soil/pkg/tool/goflightd/types/status"

func (s *Service) Status() (*status.Status, error) {
	events, e := s.store.EventCount()

	if e != nil {
		return nil, e
	}

	snapshots, f := s.store.SnapshotCount()

	if f != nil {
		return nil, f
	}

	marks, g := s.store.MarkCount()

	if g != nil {
		return nil, g
	}

	last, h := s.store.LastEventTime()

	if h != nil {
		return nil, h
	}

	return status.New(events, snapshots, marks, last), nil
}
