package service

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"

func (s *Service) UnsubscribeMergeRequest(
	instance string,
	project int64,
	request int64,
	callsign string,
) (bool, error) {
	if e := s.requireWatches(); e != nil {
		return false, e
	}

	k := watched_request.New(instance, project, request)
	count, e := s.store.Delete(callsign, k)

	if e != nil || count == 0 {
		return false, e
	}

	return true, s.forget(k)
}
