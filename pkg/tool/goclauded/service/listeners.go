package service

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
)

func (s *Service) listeners(name string) ([]*session.Session, error) {
	sessions, e := s.store.ListSessions()

	if e != nil {
		return nil, e
	}

	var result []*session.Session

	for i := range sessions {
		v := &sessions[i]

		if v.CallsignValue() == name {
			continue
		}

		value, f := s.store.GetLabel(v.Identifier, constant.BroadcastLabel)

		if f != nil {
			return nil, f
		}

		if value == constant.BroadcastOff {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
