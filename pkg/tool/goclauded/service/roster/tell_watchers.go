package roster

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"slices"
)

func (s *Service) tellWatchers(
	sessionIdentifier string,
	body string,
	d constant.Delivery,
) error {
	sessions, e := s.store.ListSessions()

	if e != nil {
		return e
	}

	for i := range sessions {
		v := &sessions[i]

		if v.Identifier == sessionIdentifier {
			continue
		}

		values, f := s.store.GetLabel(
			v.Identifier,
			constant.CompactionWatchLabel,
		)

		if f != nil {
			return f
		}

		if !slices.Contains(values, constant.CompactionWatchValue) {
			continue
		}

		if _, g := s.message.SendNotification(
			v.CallsignValue(),
			constant.Compaction,
			body,
			d,
		); g != nil {
			return g
		}
	}

	return nil
}
