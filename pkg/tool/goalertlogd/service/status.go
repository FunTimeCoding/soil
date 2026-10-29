package service

import "github.com/funtimecoding/soil/pkg/tool/goalertlogd/types/status"

func (s *Service) Status() (*status.Status, error) {
	result := status.New()
	var e error

	if result.Total, e = s.store.Count(); e != nil {
		return nil, e
	}

	var f error

	if result.Firing, f = s.store.UnresolvedCount(); f != nil {
		return nil, f
	}

	var g error

	if result.LatestIdentifier, g = s.store.LatestIdentifier(); g != nil {
		return nil, g
	}

	return result, nil
}
