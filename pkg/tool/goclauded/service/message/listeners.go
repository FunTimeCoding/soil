package message

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"
	"slices"
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

		values, f := s.store.GetLabel(v.Identifier, constant.BroadcastLabel)

		if f != nil {
			return nil, f
		}

		if slices.Contains(values, constant.BroadcastOff) {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
