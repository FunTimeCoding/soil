package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) recipients(
	name string,
	to []string,
) ([]*session.Session, error) {
	var result []*session.Session

	for _, c := range to {
		found, e := s.store.SessionByCallsign(c)

		if e != nil {
			return nil, e
		}

		if found == nil {
			return nil, not_found.New(constant.Callsign, c)
		}

		result = append(result, found)
	}

	if len(to) > 0 {
		return result, nil
	}

	return s.listeners(name)
}
